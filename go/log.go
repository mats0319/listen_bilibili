package lb

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	newLogger *log.Logger
	file      *os.File
)

func init() {
	err := openLogFile()
	if err != nil {
		waitAndExit(3)
	}

	newLogger = log.New(file, "", log.LstdFlags)
}

func Printf(format string, v ...any) {
	log.Printf(format, v...)
	newLogger.Printf(format, v...)
}

func Println(v ...any) {
	log.Println(v...)
	newLogger.Println(v...)
}

func openLogFile() error {
	// log file not init yet in this func, just write log to cmd window
	err := os.MkdirAll("./log/", 0777)
	if err != nil {
		log.Println("mkdir './log/' failed, error: ", err)
		return err
	}

	filename := fmt.Sprintf("./log/log_%d.log", time.Now().Unix())

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println("open log file failed, error: ", err)
		return err
	}

	return nil
}

func waitAndExit(second int) {
	for i := 0; i < second; i++ {
		log.Printf("Exit in %ds.\n", second-i)
		time.Sleep(time.Second)
	}

	os.Exit(-1)
}
