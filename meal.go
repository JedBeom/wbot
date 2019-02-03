package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	sm "github.com/JedBeom/schoolmeal"
)

var (
	meals []sm.Meal
)

// 급식을 불러옴
func getMeals() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovering:", r)
			time.Sleep(time.Second)
			getMeals()
		}

	}()

	school := sm.School{
		Code: "Q100005451",
		Kind: sm.Middle,
		Zone: sm.Jeonnam,
	}

	now := time.Now()

	// 토요일일 경우 다음주 급식
	if now.Weekday() == time.Saturday {
		now = now.AddDate(0, 0, 1)
	}

	// 점심대의 급식을 가져온다
	todayMeals, err := school.GetWeekMeal(sm.Timestamp(now), sm.Lunch)
	if err != nil {
		log.Println(err)
		return
	}

	meals = todayMeals

}

// 급식 스킬
func MealSkill(w http.ResponseWriter, r *http.Request) {

	// payload 파싱
	payload, err := ParsePayload(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	logger(payload)

	// 급식 스킬인데 요일이 없다면
	if payload.Weekday == "" {
		log.Println("No weekday in payload")

		w.WriteHeader(400)
		return
	}

	var simpleText string
	var weekdayCode int

	// 한글에 따라 index 번호 정하기
	switch payload.Weekday {
	case "월요일":
		weekdayCode = 1
	case "화요일":
		weekdayCode = 2
	case "수요일":
		weekdayCode = 3
	case "목요일":
		weekdayCode = 4
	case "금요일":
		weekdayCode = 5
	case "토요일", "일요일":
		simpleText = "토요일과 일요일 급식은 없어요."
	default:
		simpleText = "무슨 말인지 모르겠어요."
	}

	var meal sm.Meal
	// 뭐? 받아온 급식이 없어?
	if len(meals) != 0 {
		meal = meals[weekdayCode]
	} else {
		simpleText = "급식을 가져올 수 없어요."
	}

	// 위에서 문제가 없었다면
	if simpleText == "" {
		var content string
		if meal.Content != "" {
			// \n을 \\n으로 치환
			content = strings.Replace(meal.Content, "\n", "\\n", -1)
		} else {
			content = "급식 정보가 없어요."
		}
		simpleText = meal.Date + "\\n" + content
	}

	format := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"도움말","action":"message"},{"label":"월요일","action":"message"},{"label":"화요일","action":"message"},{"label":"수요일","action":"message"},{"label":"목요일","action":"message"},{"label":"금요일","action":"message"}]}}`

	/*
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
					},
					{
						"label": "월요일",
						"action": "message"
					},
					{
						"label": "화요일",
						"action": "message"
					},
					{
						"label": "수요일",
						"action": "message"
					},
					{
						"label": "목요일",
						"action": "message"
					},
					{
						"label": "금요일",
						"action": "message"
					}
				]
			}
		}`
	*/

	// blockId: 5c28aa155f38dd44d86a0f85

	output := fmt.Sprintf(format, simpleText)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write([]byte(output))
	if err != nil {
		log.Println("Error while w.Write:", err)
	}

}
