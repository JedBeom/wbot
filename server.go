package main

import (
	"log"
	"net/http"
)

func serve() {
	server := http.Server{
		Addr: config.Port,
	}

	http.HandleFunc("/meal", mealSkill)
	http.HandleFunc("/airq", airqSkill)
	http.HandleFunc("/dday", dDaySkill)
	http.HandleFunc("/fb_posts", fbSkill)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server Error:", err)
	}

}
