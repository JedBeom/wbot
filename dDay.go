package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

var (
	dDayT    *template.Template
	DdayText string
)

func init() {
	format := `📅 학교 주요 일정이에요!
{{ range . }}
{{ .DateString }} {{ .Name }}
D{{ .LeftDays }}
{{ end }}`

	dDayT = template.Must(template.New("format").Parse(format))
}

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
				"label": "급식",
				"action": "message"
			},
			{
				"label": "미세먼지",
				"action": "message"
			}
		]
	}
}`

	output := fmt.Sprintf(format, DdayText)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(output))
}

func GetEvents() {
	file, err := ioutil.ReadFile("events.json")
	if err != nil {
		log.Println(err)
		return
	}

	var events []Event
	err = json.Unmarshal(file, &events)

	var RealEvents []Event

	now := time.Now()
	nowMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	for _, value := range events {
		parsedDate, err := time.Parse("2006/01/02", value.DateString)
		if err != nil {
			log.Println(err)
			continue
		}

		value.Date = parsedDate

		left := value.Date.Sub(nowMidnight).Hours()
		value.LeftDays = -int(left / 24)
		RealEvents = append(RealEvents, value)
	}

	var tpl bytes.Buffer
	dDayT.Execute(&tpl, RealEvents)

	DdayText = strings.Replace(tpl.String(), "\n", "\\n", -1)

}
