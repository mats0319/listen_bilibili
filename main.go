package main

import (
	"github.com/mats9693/listenBilibili/go"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getList", listen_bilibili.OnGetList)
	http.HandleFunc("/getOriginURL", listen_bilibili.OnGetOriginURL)

	err := http.ListenAndServe("0.0.0.0:9693", nil)
	if err != nil {
		log.Println("listen and serve failed, err:", err)
	}
}
