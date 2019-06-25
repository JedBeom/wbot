package main

import (
	"fmt"
	"log"
	"net/http"
)

// 디데이 스킬
func dDaySkill(w http.ResponseWriter, r *http.Request) {
	payload, err := ParsePayload(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	logger(payload)

	format := `{
	"version": "2.0",
	"template": {
		"outputs": [
			{
				"simpleText": {
					"text": "%s"
				}
			}
		],
		"quickReplies": [
			{
				"label": "새로고침",
				"action": "block",
				"blockId": "%s"
			}
		]
	}
}`

	output := fmt.Sprintf(format, DdayText, payload.BlockID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write([]byte(output))
	if err != nil {
		log.Println("Error while writing in dDay:", err)
	}
}
