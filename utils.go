package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/JedBeom/wbot_new/model"

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
	history.Params = make(map[string]string)
	paramsJSON, _, _, _ := jsonparser.Get(payloadJSON, "action", "params")
	fmt.Println(string(paramsJSON))
	err = json.Unmarshal(paramsJSON, &history.Params)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
