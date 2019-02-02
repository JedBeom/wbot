package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JedBeom/airq"
)

var (
	// 미세먼지 저장용 전역 변수
	hangulQ HangulQ
)

func init() {
	// 인증키 가져오기
	err := airq.GetKeyFile("airq_key.txt")
	if err != nil {
		panic(err)
	}

}

// 미세먼지 불러오기
func getAirq(stationName string) {
	// init
	hangulQ = HangulQ{}

	hangulQ.Station = stationName

	// 미세먼지 가져오기
	quality, err := airq.NowByStation(stationName)
	// 문제가 있고 연향동이면
	if err != nil && stationName == "연향동" {
		getAirq("장천동") // 장천동으로 다시 가져온다
		return
		// 문제가 있고 장천동이면
	} else if err != nil && stationName == "장천동" {
		log.Println("Error while getting airq; stationName: 장천동", err)
		hangulQ.Error = err
		return
	}

	// 등급에 따라 한글 등급을 매긴다
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

	// 더 안좋은 등급을 가져온다
	if quality.Pm10GradeWHO > quality.Pm25GradeWHO {
		hangulQ.MixedRate = quality.Pm10GradeWHO
	} else {
		hangulQ.MixedRate = quality.Pm25GradeWHO
	}

	return

}

// 미세먼지 스킬
func AirqSkill(w http.ResponseWriter, r *http.Request) {
	payload, err := ParsePayload(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	logger(payload)

	var simpleText string
	// 미세먼지에 문제가 있으면
	if hangulQ.Error != nil {
		simpleText = "미세먼지 측정소가 응답하지 않아요."
		hangulQ.MixedRate = 0
	} else {

		template := "미세먼지는 %s, 초미세먼지는 %s!"
		simpleText = fmt.Sprintf(template, hangulQ.Pm10, hangulQ.Pm25)
	}

	format := `{"version":"2.0","template":{"outputs":[{"basicCard":{"title":"%s","description":"측정소: %s","thumbnail":{"imageUrl":"https://raw.githubusercontent.com/JedBeom/wbot_new/master/img/%d.jpg"}}}],"quickReplies":[{"label":"도움말","action":"message"},{"label":"새로고침","action":"block","blockId":"%s"}]}}`

	/*
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
						"label": "도움말",
						"action": "message"
					},
					{
						"label": "새로고침",
						"action": "block",
						"blockId": "%s"
					}
				]
			}
		}`
	*/

	output := fmt.Sprintf(format, simpleText, hangulQ.Station, hangulQ.MixedRate, payload.BlockID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(output))

}
