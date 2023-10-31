package lb

import (
	"github.com/mats9693/listenBilibili/api/go"
	"gopkg.in/yaml.v3"
	"os"
)

var list = &api.List{}

func LoadList(filename string) error {
	listBytes, err := os.ReadFile(filename)
	if err != nil {
		Println("read list file failed, error: ", err)
		return err
	}

	err = yaml.Unmarshal(listBytes, list)
	if err != nil {
		Println("deserialize list failed, error: ", err)
		return err
	}

	Println("> Init List Finished.")

	return nil
}
