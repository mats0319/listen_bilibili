package main

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	url := getVideoURL("https://m.bilibili.com/video/BV1nk4y1P79g?p=22")
	log.Println("> url: ", url)

	//err := http.ListenAndServe("0.0.0.0:9693", nil)
	//if err != nil {
	//	log.Println("listen and serve failed, err:", err)
	//}
}

func getVideoURL(url string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
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
