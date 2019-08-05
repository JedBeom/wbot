package main

import (
	"log"
	"net/http"

	"github.com/JedBeom/wbot_new/model"
)

func feedbackSkill(w http.ResponseWriter, r *http.Request) {
	history, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(400)
		return
	}

	if _, err := feedbackFile.WriteString(history.Params["feedback"] + "\n"); err != nil {
		log.Print("Feedback Error:", err)
	}
	_, _ = w.Write([]byte(`{"version": "2.0"}`))
	w.WriteHeader(200)
}
