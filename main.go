package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/robfig/cron.v2"
)

func midnightDo() {
	go getMeals()
	go getEvents()
	go getFBPosts()
}

func workInit() {

	c := cron.New()

	// every 12 am
	if _, err := c.AddFunc("0 0 0 * * *", midnightDo); err != nil {
		panic(err)
	}

	// Every xx:14
	if _, err := c.AddFunc("0 14 * * * *", getAirqDefault); err != nil {
		panic(err)
	}

	// init
	midnightDo()

	setAirqKey()
	getAirq("연향동")

	go c.Start()
}

func getAirqDefault() {
	getAirq("연향동")
}

func main() {
	workInit()

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./wbot_new [port]")
		os.Exit(1)
	}

	port := ":" + os.Args[1]
	server := http.Server{
		Addr: port,
	}

	accessLog, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	defer func() {
		err = accessLog.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println("Starting")
	// Set logo output
	log.SetOutput(accessLog)
	log.Println("Server Started")

	http.HandleFunc("/meal", MealSkill)
	http.HandleFunc("/airq", AirqSkill)
	http.HandleFunc("/dday", DDaySkill)
	http.HandleFunc("/fb_posts", fbSkill)

	err = server.ListenAndServe()
	if err != nil {
		log.Println("Server Error:", err)
	}

}
