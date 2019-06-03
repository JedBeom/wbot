package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JedBeom/airq"
	"github.com/pkg/errors"
)

var (
	// 미세먼지 저장용 전역 변수
	hangulQ HangulQ
)

func setAirqKey() {
	// 인증키 가져오기
	err := airq.SetKey(config.AirqKey)
	if err != nil {
		panic(err)
	}

}

// 미세먼지 불러오기
func getAirq() {

	// initial
	hangulQ = HangulQ{}

	stations := []string{
		"연향동",
		"장천동",
	}

	var quality *airq.AirQuality
	for _, station := range stations {
		if q, err := airq.NowByStation(station); err == nil {
			quality = &q
			hangulQ.Station = station
			break
		}
	}

	if quality == nil {
		hangulQ.Error = errors.New("No airq;")
		return
	}

	hangulQ.TimeString = quality.DataTimeString

	// 등급에 따라 한글 등급을 매긴다
	hangulQ.Pm10 = rateToKo(quality.Pm10GradeWHO)
	hangulQ.Pm25 = rateToKo(quality.Pm25GradeWHO)

	// 더 안좋은 등급을 가져온다
	if quality.Pm10GradeWHO > quality.Pm25GradeWHO {
		hangulQ.MixedRate = quality.Pm10GradeWHO
	} else {
		hangulQ.MixedRate = quality.Pm25GradeWHO
	}

	return

}

// 미세먼지 스킬
func airqSkill(w http.ResponseWriter, r *http.Request) {
	payload, err := ParsePayload(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	logger(payload)

	var simpleText string
	var description string

	// 미세먼지에 문제가 있으면
	if hangulQ.Error != nil {
		simpleText = "미세먼지 측정소가 응답하지 않아요."
		description = "한 시간 뒤에 다시 시도해 주세요."
		hangulQ.MixedRate = 0
	} else {

		if hangulQ.Pm10 == hangulQ.Pm25 {
			simpleText = fmt.Sprintf("미세먼지와 초미세먼지는 %s 상태!", hangulQ.Pm10)
		} else {
			simpleText = fmt.Sprintf("미세먼지는 %s, 초미세먼지는 %s!", hangulQ.Pm10, hangulQ.Pm25)
		}
		description = fmt.Sprintf("측정소: %s | 측정 시간: %s", hangulQ.Station, hangulQ.TimeString)

	}

	format := `{"version":"2.0","template":{"outputs":[{"basicCard":{"title":"%s","description":"%s","thumbnail":{"imageUrl":"https://raw.githubusercontent.com/JedBeom/wbot_new/master/img/%d.jpg"}}}],"quickReplies":[{"label":"도움말","action":"message"},{"label":"새로고침","action":"block","blockId":"%s"}]}}`

	output := fmt.Sprintf(format, simpleText, description, hangulQ.MixedRate, payload.BlockID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write([]byte(output))
	if err != nil {
		log.Println("airqSkill:", err)
	}

}

func rateToKo(value int) (rate string) {

	switch value {
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

	return
}
