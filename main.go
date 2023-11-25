package main

import (
	"github.com/mats9693/listen_bilibili/go"
	"net/http"
	"os/exec"
)

func main() {
	initList()

	openWebpage()

	err := http.ListenAndServe(":9693", lb.GetHandler())
	if err != nil {
		lb.Println("listen and serve failed, error: ", err)
	}
}

func initList() {
	err := lb.ReadList()
	if err != nil {
		lb.Println("load list failed, error: ", err)
		return
	}

	lb.Println("> Init List Finished.")
}

func openWebpage() {
	// auto open webpage in Windows OS
	err := exec.Command("cmd", "/c start http://127.0.0.1:9693").Start()
	if err != nil {
		lb.Println("auto open webpage failed：", err)
		lb.Println("please visit manually：http://127.0.0.1:9693")
	}
}
