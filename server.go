package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func serve() {

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(MiddlewareHistory)
		r.Use(middleware.AllowContentType("application/json"))

		// Original
		r.Route("/original", func(r chi.Router) {
			r.Post("/meal", SkillMeal)
			r.Post("/airq", SkillAirq)
			r.Post("/events", SkillEvents)
		})
		// New
		r.Route("/new", func(r chi.Router) {
			r.Post("/facebook", SkillFacebook)
			r.Post("/feedback", SkillFeedback)
		})
		// School Request
		r.Route("/school", func(r chi.Router) {
			r.Post("/reports", SkillReport)
			r.Post("/checker", SkillChecker)
			r.Post("/enter", SkillEnterStudentInfo)
		})
	})

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

func status(w http.ResponseWriter, r *http.Request) {
	write(w, r.UserAgent())
	return
}
