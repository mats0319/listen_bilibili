package lb

import (
	"os/exec"
)

func OpenWebpage() {
	// auto open webpage in Windows OS
	err := exec.Command("cmd", "/c start http://127.0.0.1:9693").Start()
	if err != nil {
		Println("auto open webpage failed：", err)
		Println("please visit manually：http://127.0.0.1:9693")
	}
}
