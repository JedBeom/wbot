package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func serve() {

	r := chi.NewRouter()

	r.Use(MiddlewareHistory)

	// Original
	r.Mount("/original/", RouterOriginal())
	// New
	r.Post("/new/facebook", facebookSkill)
	// School Request
	r.Post("/school/reports", nil)
	// Status Checking
	r.Get("/status", status)

	server := http.Server{
		Addr:    config.Port,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server Error:", err)
	}

}

func RouterOriginal() http.Handler {
	r := chi.NewRouter()
	r.Post("/meal", mealSkill)
	r.Post("/airq", airqSkill)
	r.Post("/events", dDaySkill)
	r.Post("/feedback", feedbackSkill)
	return r
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte(r.UserAgent()))
	return
}
