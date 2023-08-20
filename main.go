package main

import (
	"encoding/json"
	"github.com/mats9693/listenBilibili/api/go"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

var list = &api.List{}

func main() {
	listBytes, err := os.ReadFile("./listen_bilibili.json")
	if err != nil {
		log.Println("read file failed, err: ", err)
		return
	}

	err = json.Unmarshal(listBytes, list)
	if err != nil {
		log.Println("deserialize list failed, err: ", err)
	}

	http.HandleFunc("/getList", onGetList)
	http.HandleFunc("/getOriginURL", onGetOriginURL)
	err = http.ListenAndServe("0.0.0.0:9693", nil)
	if err != nil {
		log.Println("listen and serve failed, err:", err)
	}
}

// onGetList return List
func onGetList(w http.ResponseWriter, _ *http.Request) {
	res := &api.GetListRes{}

	listBytes, err := json.Marshal(list)
	if err != nil {
		log.Println("json marshal failed, err:", err)
		res.Err = err.Error()
	}

	res.List = string(listBytes)

	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("json marshal failed, err:", err)
	}

	_, err = w.Write(resBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// onGetOriginURL according to 'music id', match 'bv' and use 'bv' to get 'origin address'
func onGetOriginURL(w http.ResponseWriter, req *http.Request) {
	musicID := req.PostFormValue("music_id")

	bv := getBv(musicID)
	url := getOriginURL(bv)

	res := &api.GetOriginURLRes{URL: url}

	if len(bv) < 1 || len(url) < 1 {
		res.Err = "runs error"
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("json marshal failed, err:", err)
	}

	_, err = w.Write(resBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getBv match 'bv' according 'music id', return empty string if not matched
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

func getOriginURL(bv string) string {
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

	reRule := regexp.MustCompile(`"readyVideoUrl":\s*"([^"]*)"`)
	result := reRule.FindAllSubmatch(body, -1)
	if len(result) < 1 || len(result[0]) < 2 {
		log.Println("RE match failed.")
		return ""
	}

	return string(result[0][1])
}
