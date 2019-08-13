package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/JedBeom/wbot_new/model"
)

func SkillDice(w http.ResponseWriter, r *http.Request) {
	h, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"다시 해볼게!","action":"block", "blockId": "%s"},{"label": "%d면체","action":"message"}]}}`

	dice, err := strconv.Atoi(h.Params["dice"])
	if err != nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "지금은 6, 7, 8, 9, 10, 11, 12, 20, 30면체만 지원해.", h.BlockID, 6))
		return
	}

	random := rand.Intn(dice) + 1
	randomStr := strconv.Itoa(random)
	switch randomStr[len(randomStr)-1] {
	case '2', '4', '5', '9':
		randomStr += "가"
	default:
		randomStr += "이"
	}
	result := fmt.Sprintf("🎲 %d면체 주사위를 돌려서 %s 나왔어!", dice, randomStr)
	writeOK(w, fmt.Sprintf(simpleFormat, result, h.BlockID, dice))

}

func SkillYesOrNo(w http.ResponseWriter, r *http.Request) {
	h, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"내 대답은 %s(%d%%)."}}],"quickReplies":[{"label":"다시 물어볼게...","action":"block", "blockId": "%s"}]}}`
	percent := rand.Intn(101)
	yesOrNo := "예"
	if percent < 50 {
		yesOrNo = "아니요"
	}

	writeOK(w, fmt.Sprintf(simpleFormat, yesOrNo, percent, h.BlockID))
}
