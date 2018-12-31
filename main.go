package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jasonlvhit/gocron"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./wbot_new [port]")
		os.Exit(1)
	}

	port := ":" + os.Args[1]
	server := http.Server{
		Addr: port,
	}

	gocron.Every(1).Day().At("00:00").Do(getMeals)
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

	getMeals()
	getAirq("연향동")

	http.HandleFunc("/meal", MealSkill)
	http.HandleFunc("/airq", AirqSkill)
	server.ListenAndServe()
}
