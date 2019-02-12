package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	dDayT    *template.Template
	DdayText string
)

func init() {
	format := `📅 학교 주요 일정이에요!
{{ range . }}
{{ .DateString }} {{ .Name }}
{{if .LeftDays}}D{{ .LeftDays }}{{else}}D-DAY 🎉{{end}}
{{ end }}`

	dDayT = template.Must(template.New("format").Parse(format))
}

// 디데이 스킬
func DDaySkill(w http.ResponseWriter, r *http.Request) {
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
				"label": "도움말",
				"action": "message"
			}
		]
	}
}`

	output := fmt.Sprintf(format, DdayText)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write([]byte(output))
	if err != nil {
		log.Println("Error while writing in dDay:", err)
	}
}
