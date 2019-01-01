package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JedBeom/airq"
)

var (
	hangulQ HangulQ
)

// 미세먼지 불러오기
func getAirq(stationName string) {
	err := airq.LoadServiceKey("airq_key.txt")
	if err != nil {
		log.Println(err)
	}

	hangulQ.Station = stationName

	quality, err := airq.GetAirqOfNowByStation(stationName)
	if err != nil && stationName == "연향동" {
		getAirq("장천동")
		return
	} else if err != nil && stationName == "장천동" {
		hangulQ.Error = err
		return
	}

	var rate string
	switch quality.Pm10GradeWHO {
	case 1:
		rate = "최고"
	case 2:
		rate = "좋음"
	case 3:
		rate = "양호"
	case 4:
		rate = "보통"
	case 5:
		rate = "나쁨"
	case 6:
		rate = "상당히 나쁨"
	case 7:
		rate = "매우 나쁨"
	case 8:
		rate = "최악"
	}
	hangulQ.Pm10 = rate

	switch quality.Pm25GradeWHO {
	case 1:
		rate = "최고"
	case 2:
		rate = "좋음"
	case 3:
		rate = "양호"
	case 4:
		rate = "보통"
	case 5:
		rate = "나쁨"
	case 6:
		rate = "상당히 나쁨"
	case 7:
		rate = "매우 나쁨"
	case 8:
		rate = "최악"
	}
	hangulQ.Pm25 = rate

	if quality.Pm10GradeWHO > quality.Pm25GradeWHO {
		hangulQ.MixedRate = quality.Pm10GradeWHO
	} else {
		hangulQ.MixedRate = quality.Pm25GradeWHO
	}

	return

}

func AirqSkill(w http.ResponseWriter, r *http.Request) {
	payload, err := ParsePayload(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	logger(payload)

	var simpleText string
	if hangulQ.Error != nil {
		simpleText = "미세먼지 측정소가 응답하지 않아요."
	} else {

		template := "미세먼지는 %s, 초미세먼지는 %s!"
		simpleText = fmt.Sprintf(template, hangulQ.Pm10, hangulQ.Pm25)
	}

	format := `{
	"version": "2.0",
	"template": {
		"outputs": [
			{
				"basicCard": {
					"title": "%s",
					"description": "측정소: %s",
					"thumbnail": {
						"imageUrl": "https://raw.githubusercontent.com/JedBeom/wbot_new/master/img/%d.jpg"
					}
				}
			}
		],
		"quickReplies": [
			{
				"label": "급식",
				"action": "message"
			}
		]
	}
}`

	output := fmt.Sprintf(format, simpleText, hangulQ.Station, hangulQ.MixedRate)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(output))

}
