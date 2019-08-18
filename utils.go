package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JedBeom/wbot/model"

	"github.com/buger/jsonparser"
)

// Parse payload from json
func ParseHistory(body io.Reader) (history model.History, err error) {

	payloadJSON, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		return
	}

	history.BlockName, _ = jsonparser.GetString(payloadJSON, "userRequest", "block", "name")
	history.BlockID, _ = jsonparser.GetString(payloadJSON, "userRequest", "block", "id")
	history.UserID, _ = jsonparser.GetString(payloadJSON, "userRequest", "user", "id")
	history.Utterance, _ = jsonparser.GetString(payloadJSON, "userRequest", "utterance")

	// Only for SkillCancelReport
	history.ContextDetail, _ = jsonparser.GetString(payloadJSON, "contexts", "[0]", "params", "detail", "value")

	history.Params = make(map[string]string)
	paramsJSON, _, _, _ := jsonparser.Get(payloadJSON, "action", "params")
	err = json.Unmarshal(paramsJSON, &history.Params)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func paramsToReport(h model.History, studentID int) (report model.Report, err error) {
	report = model.Report{
		ReportType:      h.Params["report_type"],
		UserID:          h.UserID,
		HistoryID:       h.ID,
		What:            h.Params["what"],
		Detail:          h.Params["detail"],
		TargetStudentID: studentID,
	}

	whenStruct := struct {
		Value string `json:"value"`
	}{}

	err = json.Unmarshal([]byte(h.Params["when"]), &whenStruct)
	if err != nil {
		return
	}

	parse, err := time.Parse("2006-01-02T15:04:05", whenStruct.Value)
	if err != nil {
		return
	}
	seoul, _ := time.LoadLocation("Asia/Seoul")
	report.When = parse.In(seoul).Add(-9 * time.Hour)

	return
}

func extractInt(text string) int {
	if len(text) > 2 {
		return int(text[0] - '0')
	}

	return 0
}

func writeOK(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(text))
}
