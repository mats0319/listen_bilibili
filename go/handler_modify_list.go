package lb

import (
	"encoding/json"
	"github.com/mats9693/listen_bilibili/api/go"
	"net/http"
)

func onModifyList(req *http.Request) ([]byte, error) {
	res := &api.ModifyListRes{}

	listStr := req.PostFormValue("list")

	// backup current list file
	err := backupList()
	if err != nil {
		Println("backup list failed, error: ", err)
		res.Err = err.Error()
		return nil, err
	}

	// write 'list str' into list file
	err = writeList(listStr)
	if err != nil {
		Println("modify list file failed, error: ", err)
		res.Err = err.Error()
		return nil, err
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		Println("json marshal failed, error:", err)
		res.Err = err.Error()
		return nil, err
	}

	return resBytes, nil
}
