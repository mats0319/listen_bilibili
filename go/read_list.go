package listen_bilibili

import (
	"encoding/json"
	"github.com/mats9693/listenBilibili/api/go"
	"log"
	"os"
)

var list = &api.List{}

func init() {
	listBytes, err := os.ReadFile("./listen_bilibili.json")
	if err != nil {
		log.Println("read file failed, err: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal(listBytes, list)
	if err != nil {
		log.Println("deserialize list failed, err: ", err)
		os.Exit(1)
	}

	log.Println("> Init List Finished.")
}
