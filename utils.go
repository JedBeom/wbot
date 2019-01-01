package main

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
)

// Write log
func logger(payload Payload) {
	log.Printf("%s %s %s", payload.BlockName, payload.UserID, payload.Utterance)
}

// Parse payload from json
func ParsePayload(body io.Reader) (payload Payload, err error) {

	payloadJSON, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		return
	}

	payload.BlockName, _ = jsonparser.GetString(payloadJSON, "userRequest", "block", "name")
	payload.BlockID, _ = jsonparser.GetString(payloadJSON, "userRequest", "block", "id")
	payload.UserID, _ = jsonparser.GetString(payloadJSON, "userRequest", "user", "id")
	payload.Utterance, _ = jsonparser.GetString(payloadJSON, "userRequest", "utterance")
	payload.Weekday, _ = jsonparser.GetString(payloadJSON, "action", "detailParams", "요일", "value")
	return
}
