package main

import (
	"log"
	"os"
)

func main() {
	loadConfig("config.json")
	workInit()

	errorLog, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
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
