package listen_bilibili

import (
	"encoding/json"
	"github.com/mats9693/listenBilibili/api/go"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	originURLRE      = `"readyVideoUrl":\s*"([^"]*)"`
	internalErrorMsg = "Internal Server Error"
)

func BindHTMLFile(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	path := dir + "/ui/dist" + r.RequestURI
	if r.RequestURI == "/" {
		path += "index.html"
	}
	http.ServeFile(w, r, path)
}

// OnGetList return List
func OnGetList(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res := &api.GetListRes{}

	listBytes, err := json.Marshal(list)
	if err != nil {
		log.Println("json marshal failed, err:", err)
		res.Err = err.Error()
	} else {
		res.List = string(listBytes)
	}

	responseHTTP(w, res)
}

// OnGetOriginURL according to 'music id', match 'bv' and use 'bv' to get 'origin address'
func OnGetOriginURL(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	res := &api.GetOriginURLRes{}

	if req.Method == http.MethodOptions {
		responseHTTP(w, res)
		return
	}

	musicID := req.PostFormValue("music_id")

	url := getOriginURL(musicID)
	if len(url) < 1 {
		log.Println("get origin url failed")
		res.Err = internalErrorMsg
	} else {
		res.URL = url
	}

	responseHTTP(w, res)
}

func responseHTTP(w http.ResponseWriter, res any) {
	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("json marshal failed, err:", err)
	}

	_, err = w.Write(resBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getOriginURL(musicID string) string {
	bv := getBv(musicID)
	if len(bv) < 1 {
		log.Println("unmatched music id: ", musicID)
		return ""
	}

	// simulate mobile invoke, get HTML file, 'origin url' is in it
	client := &http.Client{}

	req, err := http.NewRequest("GET", bv, nil)
	if err != nil {
		log.Println("create new request failed, err:", err)
		return ""
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Mobile Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		log.Println("http get failed, err:", err)
		return ""
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("read http res body failed, err:", err)
		return ""
	}

	// get 'origin url' use RE
	originURL := matchOriginURL(body)
	if len(originURL) < 1 {
		errMsg := "RE match failed"
		log.Println(errMsg)
		return ""
	}

	return originURL
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

func matchOriginURL(data []byte) string {
	reRule := regexp.MustCompile(originURLRE)
	result := reRule.FindAllSubmatch(data, -1)
	if len(result) < 1 || len(result[0]) < 2 {
		log.Println("RE match failed.")
		return ""
	}

	return string(result[0][1])
}
