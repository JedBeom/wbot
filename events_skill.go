package main

import (
	"fmt"
	"net/http"

	"github.com/JedBeom/wbot/model"
)

// 디데이 스킬
func SkillEvents(w http.ResponseWriter, r *http.Request) {
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

	output := fmt.Sprintf(format, EventResponse, history.BlockID)
	writeOK(w, output)
}
