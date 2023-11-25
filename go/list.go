package lb

import (
	"encoding/json"
	"fmt"
	"github.com/mats9693/listen_bilibili/api/go"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

var list = &api.List{}

func ReadList() error {
	listBytes, err := os.ReadFile("./list.yaml")
	if err != nil {
		Println("read list file failed, error: ", err)
		return err
	}

	err = yaml.Unmarshal(listBytes, list)
	if err != nil {
		Println("deserialize list failed, error: ", err)
		return err
	}

	return nil
}

func writeList(listStr string) error {
	listIns := &api.List{}

	err := json.Unmarshal([]byte(listStr), listIns)
	if err != nil {
		Println("deserialize 'list' from json string failed, error: ", err)
		return err
	}

	listBytes, err := yaml.Marshal(listIns)
	if err != nil {
		Println("serialize 'list' to yaml string failed, error: ", err)
		return err
	}

	err = os.WriteFile("./list.yaml", listBytes, 0777)
	if err != nil {
		Println("modify list file failed, error: ", err)
		return err
	}

	return nil
}

func backupList() error {
	err := os.MkdirAll("./backup/", 0777)
	if err != nil {
		Println("mkdir './backup/' failed, error: ", err)
		return err
	}

	// copy 'list' file to 'backup' folder and add timestamp in filename
	listBytes, err := os.ReadFile("./list.yaml")
	if err != nil {
		Println("read list file failed, error: ", err)
		return err
	}

	filename := fmt.Sprintf("./backup/list_%d.yaml", time.Now().Unix())
	err = os.WriteFile(filename, listBytes, 0777)
	if err != nil {
		Println("backup list file failed, error: ", err)
		return err
	}

	return nil
}
