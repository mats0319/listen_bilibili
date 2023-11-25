package lb

import (
	"encoding/json"
	"github.com/mats9693/listen_bilibili/api/go"
	"net/http"
	"strconv"
)

func onGetList(req *http.Request) ([]byte, error) {
	res := &api.GetListRes{}

	// if re-load list from list file
	needReloadStr := req.PostFormValue("reload_list")
	needReload, err := strconv.ParseBool(needReloadStr)
	if err != nil {
		Printf("invalid params: %s, error: %v\n", needReloadStr, err)
		res.Err = err.Error()
		return nil, err
	}
	if needReload {
		err = ReadList()
		if err != nil {
			Println("load list failed, error: ", err)
			res.Err = err.Error()
			return nil, err
		}
	}

	res.List = *list

	resBytes, err := json.Marshal(res)
	if err != nil {
		Println("json marshal failed, error:", err)
		res.Err = err.Error()
		return nil, err
	}

	return resBytes, nil
}
