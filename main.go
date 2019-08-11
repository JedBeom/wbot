package main

import (
	"log"
	"os"
)

func main() {
	loadConfig("config.json")
	workInit()

	// Mkdir "log" if not exists
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err = os.Mkdir("log", 0777)
		if err != nil {
			panic(err)
		}
	}

	errorLog, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = errorLog.Close()
	}()

	ConnectDB()

	// Set log output
	log.SetOutput(errorLog)

	serve()
}
