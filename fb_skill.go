package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func fbSkill(w http.ResponseWriter, r *http.Request) {

	payload, err := ParsePayload(r.Body)
	if err != nil {
		log.Println("payload parsing:", err)
		return
	}
	logger(payload)

	carousel := struct {
		Items []BasicCard `json:"items"`
	}{
		Items: posts,
	}

	b, err := json.Marshal(&carousel)
	if err != nil {
		log.Println(err)
	}

	output := `{"version": "2.0","template": {"outputs":[{"simpleText": {"text": "학생회 페이스북의 최신 게시물이에요!"}},{"carousel":{"type": "basicCard", ` + string(b)[1:] + `}]}}`

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write([]byte(output))
	if err != nil {
		log.Println("Error while w.Write:", err)
	}
}
