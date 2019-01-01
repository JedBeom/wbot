package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jasonlvhit/gocron"
)

func init() {

	// every 12 am
	gocron.Every(1).Day().At("00:00").Do(getMeals)

	// Every xx:14
FinedustLoop:
	for x := 0; x < 3; x++ {

		for y := 0; y < 10; y++ {
			if x == 2 && y > 3 {
				break FinedustLoop
			}
			time := fmt.Sprintf("%d%d:16", x, y)
			gocron.Every(1).Day().At(time).Do(getAirq, "연향동")

		}
	}

	// init
	getMeals()
	getAirq("연향동")

}

func main() {
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
	defer accessLog.Close()

	log.Println("Starting")
	log.SetOutput(accessLog)
	log.Println("Server Started")

	files := http.FileServer(http.Dir("./img"))
	http.Handle("/img/", http.StripPrefix("/img/", files))

	http.HandleFunc("/meal", MealSkill)
	http.HandleFunc("/airq", AirqSkill)
	server.ListenAndServe()
}
