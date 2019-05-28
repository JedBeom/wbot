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
	http.HandleFunc("/status", status)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server Error:", err)
	}

}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("NEVER-END-IDOL"))
	return
}
