package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JedBeom/wbot_new/model"
)

func SkillReport(w http.ResponseWriter, r *http.Request) {
	history, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"다시 입력하기","action":"block", "blockId": "%s"}]}}`

	s := model.Student{
		Grade: extractInt(history.Params["grade"]),
		Class: extractInt(history.Params["class"]),
		Name:  history.Params["name"],
	}

	err := s.Find(db)
	if err != nil {
		output := fmt.Sprintf(simpleFormat, "대상 학생이 유효하지 않습니다.", history.BlockID)
		writeOK(w, output)
		return
	}

	report, err := paramsToReport(history, s.ID)
	if err != nil {
		output := fmt.Sprintf(simpleFormat,
			"예기치 못한 오류가 발생하였습니다. 잠시 후에 다시 시도해 주시기 바랍니다.", history.BlockID)
		writeOK(w, output)
		return
	}

	err = report.Create(db)
	if err != nil {
		output := fmt.Sprintf(simpleFormat,
			"왕운봇이 무언가를 실수했어요... 잠시 후에 다시 시도해 주세요.", history.BlockID)
		writeOK(w, output)
		return
	}

	report.TargetStudent = &s
	//validatorFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"다시 입력하기","action":"block", "blockId": "%s"}]}}`
	checkerBlockID := "5d50ef07ffa748000110ee9f"

	if report.ReportType == "도움필요" && history.User.StudentID == 0 {
		enterBlockID := "5d511a23ffa748000110f0f2"
		output := fmt.Sprintf(
			`{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"신고 취소하기","action":"block","blockId":"%s"},{"label":"내 정보 입력하기","action":"block","blockId":"%s"}]}}`,
			report.String()+"\\n\\n위 정보로 전송이 완료되었습니다. 내 정보가 입력되지 않았습니다. 내 정보를 입력하시면 확실한 도움을 받으실 수 있습니다.",
			checkerBlockID, enterBlockID)
		writeOK(w, output)
	} else {
		validatorFormat := `{"version":"2.0","context":{"values":[{"name":"checker","lifeSpan":5,"params":{"report_id":"%d"}}]},"template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"신고 취소하기","action":"block","blockId":"%s"}]}}`
		output := fmt.Sprintf(validatorFormat, report.ID, report.String()+"\\n\\n위 정보로 신고가 완료되었습니다. 허위 신고는 자신에게 불이익이 있습니다.", checkerBlockID)
		writeOK(w, output)
	}

	return

}

func SkillCancelReport(w http.ResponseWriter, r *http.Request) {
	history, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}
	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}]}}`

	report, err := model.FindReportByDetailAndUserID(db, history.UserID, history.ContextDetail)
	if err != nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "죄송해요, 잘 못 들었어요."))
		return
	}

	report.WasCanceled = true
	err = report.Cancel(db)
	if err != nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "죄송해요, 잘 못 들었어요."))
		return
	}

	writeOK(w, fmt.Sprintf(simpleFormat, "취소가 완료되었습니다."))
}

func SkillEnterStudentInfo(w http.ResponseWriter, r *http.Request) {
	history, ok := r.Context().Value("history").(model.History)
	if !ok {
		w.WriteHeader(500)
		return
	}

	simpleFormatWithoutRetry := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}]}}`
	if history.User.StudentID != 0 {
		writeOK(w, fmt.Sprintf(simpleFormatWithoutRetry, "이미 내 정보가 등록되어 있습니다. 이 정보는 바꿀 수 없습니다."))
		return
	}

	BlockID := "5d511a23ffa748000110f0f2"
	s := model.Student{
		Grade: extractInt(history.Params["grade"]),
		Class: extractInt(history.Params["class"]),
		Name:  history.Params["name"],
	}

	simpleFormat := `{"version":"2.0","template":{"outputs":[{"simpleText":{"text":"%s"}}],"quickReplies":[{"label":"다시 입력하기","action":"block", "blockId": "%s"}]}}`

	i, err := strconv.Atoi(history.Params["number"])
	if err != nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "유효한 학생 정보가 아닙니다. 올바르게 입력했는지 확인해주세요. 번호를 입력 시 '23번'이 아닌 '23'으로 입력해야합니다.", BlockID))
		return
	}

	s.Number = i

	err = s.Find(db)
	if err != nil {
		writeOK(w, fmt.Sprintf(simpleFormat, "유효한 학생 정보가 아닙니다. 올바르게 입력했는지 확인해주세요.", BlockID))
		return
	}

	history.User.StudentID = s.ID
	err = history.User.Update(db)
	if err != nil {
		log.Println(err)
		writeOK(w, fmt.Sprintf(simpleFormat, "유효한 학생 정보가 아닙니다. 올바르게 입력했는지 확인해주세요.", BlockID))
		return
	}
	writeOK(w, fmt.Sprintf(simpleFormatWithoutRetry, "성공적으로 내 정보를 등록하였습니다."))
}
