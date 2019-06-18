package main

import (
	"log"
	"net/http"
)

func FeedBack(w http.ResponseWriter, r *http.Request) {
	payload, err := ParsePayload(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if _, err := feedbackFile.WriteString(payload.NormalText + "\n"); err != nil {
		log.Print("Feedback Error:", err)
	}
	_, _ = w.Write([]byte(`{"version": "2.0"}`))
	w.WriteHeader(200)
}
