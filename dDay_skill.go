package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JedBeom/wbot_new/model"
)

// 디데이 스킬
func dDaySkill(w http.ResponseWriter, r *http.Request) {
	history, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(400)
		return
	}

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

	output := fmt.Sprintf(format, DdayText, history.BlockID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(output))
	if err != nil {
		log.Println("Error while writing in dDay:", err)
	}
}
