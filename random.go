package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/JedBeom/wbot/model"
)

func SkillDice(w http.ResponseWriter, r *http.Request) {
	h, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"다른 주사위 던질래!","action":"block", "blockId": "%s"},{"label": "%d면체","action":"message"}]}}`

	dice, err := strconv.Atoi(h.Params["dice"])
	if err != nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "🤔 지금은 6, 7, 8, 9, 10, 11, 12, 20, 30면체만 지원해.", h.BlockID, 6))
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

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"☯️ 내 대답은 %s(%d%%)"}}],"quickReplies":[{"label":"다시 물어볼게...","action":"block", "blockId": "%s"}]}}`
	percent := rand.Intn(101)
	yesOrNo := "예"
	if percent < 50 {
		yesOrNo = "아니요"
	}

	writeOK(w, fmt.Sprintf(simpleFormat, yesOrNo, percent, h.BlockID))
}

var (
	criteria = []string{
		" vs ",
		" ㄷ ",
		" ",
	}
)

func SkillChoice(w http.ResponseWriter, r *http.Request) {
	h, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}

	text := h.Params["text"]

	var options []string

	for _, value := range criteria {
		if strings.Contains(text, value) {
			options = strings.Split(text, value)
			break
		}
	}

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"%s","action":"block", "blockId": "%s"}]}}`
	// 선택지가 1개 이하라면 끝내기
	if len(options) < 2 || options == nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "🤔 선택지가 하나 이하인 것 같아. 두 개 이상으로 해줄래?", "알았어...", h.BlockID))
		return
	}

	// 0부터 슬라이스의 길이 중의 정수 중에서 하나를 무작위로 뽑은 다음
	// 그 숫자를 인덱스 번호로 넣어 값을 받아옴
	truth := options[rand.Intn(len(options))]
	writeOK(w, fmt.Sprintf(simpleFormat, "🔮 내 선택은 \\\""+truth+"\\\"", "다시 다시!", h.BlockID))
}
