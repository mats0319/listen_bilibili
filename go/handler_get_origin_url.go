package lb

import (
	"encoding/json"
	"errors"
	"github.com/mats9693/listen_bilibili/api/go"
	"io"
	"net/http"
	"regexp"
)

// onGetOriginURL according to 'music id', match 'bv' and use 'bv' to get 'origin address'
func onGetOriginURL(req *http.Request) ([]byte, error) {
	res := &api.GetOriginURLRes{}

	musicID := req.PostFormValue("music_id")

	musicItem := getMusic(musicID)
	if len(musicItem.ID) < 1 {
		Println("unmatched music id: ", musicID)
		res.Err = "unmatched music id: " + musicID
		return nil, errors.New("unmatched music id: " + musicID)
	}

	url, err := getOriginURL(musicItem.Bv)
	if err != nil {
		Println("get origin url failed")
		res.Err = err.Error()
	} else {
		res.URL = url
		res.Name = musicItem.Name
		res.Volume = musicItem.Volume
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		Println("json marshal failed, error:", err)
		res.Err = err.Error()
		return nil, err
	}

	return resBytes, nil
}

// getMusic return 'Music' according 'music id', return empty structure if not matched
func getMusic(musicID string) *api.Music {
	res := &api.Music{}

ALL:
	for i := range list.Playlists {
		playlistItem := list.Playlists[i]

		for j := range playlistItem.MusicList {
			musicItem := playlistItem.MusicList[j]

			if musicItem.ID == musicID {
				res = &musicItem
				break ALL
			}
		}
	}

	return res
}

func getOriginURL(bv string) (string, error) {
	data, err := getHTML(bv)
	if err != nil {
		Println("get html failed: ", err)
		return "", err
	}

	// get 'origin url' use RE
	originURL := matchOriginURL(data)
	if len(originURL) < 1 {
		Println("RE match failed")
		return "", errors.New("RE match failed")
	}

	return originURL, nil
}

// getHTML simulate mobile invoke, get HTML file, 'origin url' is in it
func getHTML(bv string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", bv, nil)
	if err != nil {
		Println("create new request failed, error:", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Mobile Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		Println("http get failed, error:", err)
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		Println("read http res body failed, error:", err)
		return nil, err
	}

	return body, nil
}

func matchOriginURL(data []byte) string {
	reRule := regexp.MustCompile(`"readyVideoUrl":\s*"([^"]*)"`)
	result := reRule.FindAllSubmatch(data, -1)
	if len(result) < 1 || len(result[0]) < 2 {
		Println("RE match failed.")
		return ""
	}

	return string(result[0][1])
}
