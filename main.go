package main

import (
	"fmt"
	"github.com/mats9693/listenBilibili/go"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	openWebpage()

	err := http.ListenAndServe(":9693", listen_bilibili.HandlersIns)
	if err != nil {
		log.Println("listen and serve failed, err:", err)
	}
}

func openWebpage() {
	// auto open webpage in Windows OS
	err := exec.Command("cmd", "/c start http://127.0.0.1:9693").Start()
	if err != nil {
		fmt.Println("auto open webpage failed：", err)
		fmt.Println("please visit manually：http://127.0.0.1:9693")
	}
}
