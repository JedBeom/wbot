package main

import (
	"log"
	"net/http"
)

func serve() {
	server := http.Server{
		Addr: config.Port,
	}

	http.HandleFunc("/meal", MealSkill)
	http.HandleFunc("/airq", AirqSkill)
	http.HandleFunc("/dday", DDaySkill)
	http.HandleFunc("/fb_posts", fbSkill)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server Error:", err)
	}

}
