package main

import (
	"log"
	"os"
)

var (
	feedbackFile *os.File
)

func main() {
	loadConfig("config.json")
	workInit()

	accessLog, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	feedbackFile, err = os.OpenFile("feedback.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = accessLog.Close()
		_ = feedbackFile.Close()
	}()

	ConnectDB()

	// Set log output
	log.SetOutput(accessLog)
	log.Println("Server Started")

	serve()
}
