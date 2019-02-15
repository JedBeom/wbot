package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func getEvents() {
	// events.json 파일 가져오기
	file, err := ioutil.ReadFile("events.json")
	if err != nil {
		log.Println(err)
		return
	}

	var events []Event
	// json 해독
	err = json.Unmarshal(file, &events)
	if err != nil {
		log.Println("Error while unmarshal events.json:", err)
		return
	}

	var RealEvents []Event

	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	for _, value := range events {
		// yyyy-mm-dd에서 time.Time 파싱
		parsedDate, err := time.Parse("2006/01/02", value.DateString)
		if err != nil {
			log.Println(err)
			continue
		}

		value.Date = parsedDate.Local().Add(time.Hour * -9)

		// 지금 마이너스 그날
		left := value.Date.Sub(midnight).Hours()
		if left < 0 {
			continue
		}
		value.LeftDays = -int(left / 24)

		RealEvents = append(RealEvents, value)
	}

	var tpl bytes.Buffer
	err = dDayT.Execute(&tpl, RealEvents)
	if err != nil {
		log.Println("Error while executing dday get...:", err)
		return
	}

	DdayText = strings.Replace(tpl.String(), "\n", "\\n", -1)

}
