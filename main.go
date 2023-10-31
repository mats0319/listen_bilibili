package main

import (
	"flag"
	"github.com/mats9693/listenBilibili/go"
	"net/http"
	"os"
)

var (
	help     bool
	listFile string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&listFile, "l", "./list.yaml", "list file")

	flag.Parse()

	if help {
		lb.FlagPrintDefaults()
		os.Exit(0)
	}
}

func main() {
	err := lb.LoadList(listFile)
	if err != nil {
		lb.Println("load list failed, error: ", err)
		return
	}

	lb.OpenWebpage()

	err = http.ListenAndServe(":9693", lb.GetHandler())
	if err != nil {
		lb.Println("listen and serve failed, error: ", err)
	}
}
