package listen_bilibili

import (
	"github.com/mats9693/listenBilibili/api/go"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

var list = &api.List{}

func init() {
	listBytes, err := os.ReadFile("./listen_bilibili.yaml")
	if err != nil {
		log.Println("read file failed, err: ", err)
		waitAndExit()
	}

	err = yaml.Unmarshal(listBytes, list)
	if err != nil {
		log.Println("deserialize list failed, err: ", err)
		waitAndExit()
	}

	log.Println("> Init List Finished.")
}

func waitAndExit() {
	time.Sleep(time.Second * 3)
	os.Exit(1)
}
