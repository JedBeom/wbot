package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	sm "github.com/JedBeom/schoolmeal"
	"github.com/buger/jsonparser"
)

var (
	meals []sm.Meal
)

// 급식을 불러옴
func getMeals() {
	school := sm.School{
		SchoolCode:     "Q100005451",
		SchoolKindCode: sm.Middle,
		Zone:           sm.Jeonnam,
	}

	now := time.Now()

	if now.Weekday() == time.Saturday {
		now = now.AddDate(0, 0, 1)
	}

	todayMeals, err := school.GetWeekMeal(sm.Timestamp(now), sm.Lunch)
	if err != nil {
		log.Println(err)
		return
	}

	meals = todayMeals

}

func MealSkill(w http.ResponseWriter, r *http.Request) {

	payloadJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)

		w.WriteHeader(400)
		return
	}

	weekday, err := jsonparser.GetString(payloadJSON, "action", "detailParams", "요일", "value")
	if err != nil {
		log.Println(err)

		w.WriteHeader(400)
		return
	}

	var simpleText string
	var weekdayCode int

	switch weekday {
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
	if len(meals) != 0 {
		meal = meals[weekdayCode]
	} else {
		simpleText = "급식 정보가 없어요."
	}

	if simpleText == "" {
		escapedContent := strings.Replace(meal.Content, "\n", "\\n", -1)
		simpleText = meal.Date + "\\n" + escapedContent
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

	// blockId: 5c28aa155f38dd44d86a0f85

	output := fmt.Sprintf(format, simpleText)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(output))

}
