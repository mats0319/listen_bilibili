package listen_bilibili

import (
	"encoding/json"
	"errors"
	"github.com/mats9693/listenBilibili/api/go"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
)

type handlerFunc = func(r *http.Request) ([]byte, error)

type Handlers struct {
	handlersMap sync.Map // request uri - func(*http.request) ([]byte, error)
}

var HandlersIns = &Handlers{}

func init() {
	HandlersIns.HandleFunc(api.URI_getList, onGetList)
	HandlersIns.HandleFunc(api.URI_getOriginalURL, onGetOriginURL)

	log.Println("> Init HTTP Handlers Finished.")
}

func (h *Handlers) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// allow cross-origin
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	// bind html file
	if request.RequestURI == "/" {
		dir, _ := os.Getwd()
		path := dir + "/ui/dist/index.html"
		http.ServeFile(writer, request, path)
		return
	}

	// skip 'http options' request
	if request.Method == http.MethodOptions {
		response(writer, []byte(""))
		return
	}

	// log req
	log.Printf("> Receive new request. uri: %s\n", request.RequestURI)

	// invoke handleFunc func
	var res []byte
	v, ok := h.handlersMap.Load(request.RequestURI)
	if !ok { // unknown uri, regard as web source
		dir, _ := os.Getwd()
		path := dir + "/ui/dist" + request.RequestURI
		if request.RequestURI == "/" {
			path += "index.html"
		}
		http.ServeFile(writer, request, path)
		return
	}

	handleFuncIns, ok := v.(handlerFunc)
	if !ok {
		http.Error(writer, "type assert error", http.StatusInternalServerError)
		return
	}

	res, err := handleFuncIns(request)
	if err != nil {
		res = []byte(err.Error())
	}

	// log res
	log.Printf("> Handle request %s: %t\n", request.RequestURI, err == nil)

	// response
	response(writer, res)
}

func (h *Handlers) HandleFunc(pattern string, hf handlerFunc) {
	log.Println("> register http handler on uri: ", pattern)

	h.handlersMap.Store(pattern, hf)
}

func response(writer http.ResponseWriter, data []byte) {
	_, err := writer.Write(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func onGetList(_ *http.Request) ([]byte, error) {
	res := &api.GetListRes{}

	listBytes, err := json.Marshal(list)
	if err != nil {
		log.Println("json marshal failed, err:", err)
		res.Err = err.Error()
	} else {
		res.List = string(listBytes)
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("json marshal failed, err:", err)
		res.Err = err.Error()
		return nil, err
	}

	return resBytes, nil
}

// onGetOriginURL according to 'music id', match 'bv' and use 'bv' to get 'origin address'
func onGetOriginURL(req *http.Request) ([]byte, error) {
	res := &api.GetOriginURLRes{}

	musicID := req.PostFormValue("music_id")

	url, err := getOriginURL(musicID)
	if err != nil {
		log.Println("get origin url failed")
		res.Err = err.Error()
	} else {
		res.URL = url
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("json marshal failed, err:", err)
		res.Err = err.Error()
		return nil, err
	}

	return resBytes, nil
}

func getOriginURL(musicID string) (string, error) {
	bv := getBv(musicID)
	if len(bv) < 1 {
		log.Println("unmatched music id: ", musicID)
		return "", errors.New("unmatched music id: " + musicID)
	}

	data, err := getHTML(bv)
	if err != nil {
		log.Println("get html failed: ", err)
		return "", err
	}

	// get 'origin url' use RE
	originURL := matchOriginURL(data)
	if len(originURL) < 1 {
		log.Println("RE match failed")
		return "", errors.New("RE match failed")
	}

	return originURL, nil
}

// getBv return 'bv' according 'music id', return empty string if not matched
func getBv(musicID string) string {
	bv := ""

ALL:
	for i := range list.Playlists {
		playlistItem := list.Playlists[i]

		for j := range playlistItem.MusicList {
			musicItem := playlistItem.MusicList[j]

			if musicItem.ID == musicID {
				bv = musicItem.Bv
				break ALL
			}
		}
	}

	return bv
}

// getHTML simulate mobile invoke, get HTML file, 'origin url' is in it
func getHTML(bv string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", bv, nil)
	if err != nil {
		log.Println("create new request failed, err:", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Mobile Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		log.Println("http get failed, err:", err)
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("read http res body failed, err:", err)
		return nil, err
	}

	return body, nil
}

func matchOriginURL(data []byte) string {
	reRule := regexp.MustCompile(`"readyVideoUrl":\s*"([^"]*)"`)
	result := reRule.FindAllSubmatch(data, -1)
	if len(result) < 1 || len(result[0]) < 2 {
		log.Println("RE match failed.")
		return ""
	}

	return string(result[0][1])
}
