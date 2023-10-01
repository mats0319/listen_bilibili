package listen_bilibili

import (
	"encoding/json"
	"errors"
	"github.com/mats9693/listenBilibili/api/go"
	"io"
	"log"
	"net/http"
	"regexp"
)

func init() {
	handlerIns.HandleFunc(api.URI_GetList, onGetList)
	handlerIns.HandleFunc(api.URI_GetOriginURL, onGetOriginURL)

	log.Println("> Init HTTP Handler Finished.")
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
