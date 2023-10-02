package lb

import (
	"flag"
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

func FlagPrintDefaults() {
	Println("Options: ")
	flag.PrintDefaults()
	flag.CommandLine.SetOutput(file)
	flag.PrintDefaults()
}

func openLogFile() error {
	var err error
	file, err = os.OpenFile("./log.log", os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		Println("open log file failed, err: ", err)
		return err
	}

	return nil
}

func waitAndExit(second int) {
	for i := 0; i < second; i++ {
		Printf("Exit in %ds.\n", second-i)
		time.Sleep(time.Second)
	}

	os.Exit(-1)
}
