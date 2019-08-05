package main

import (
	"fmt"
	"log"
	"net/http"
)

// 디데이 스킬
func dDaySkill(w http.ResponseWriter, _ *http.Request) {

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
				"label": "도움말",
				"action": "message"
			}
		]
	}
}`

	output := fmt.Sprintf(format, DdayText)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(output))
	if err != nil {
		log.Println("Error while writing in dDay:", err)
	}
}
