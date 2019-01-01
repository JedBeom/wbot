package main

import "time"

type Payload struct {
	// Request
	UserID    string
	Utterance string

	// Block
	BlockName string
	BlockID   string

	// DetailParams
	Weekday string
}

type HangulQ struct {
	Pm10 string
	Pm25 string

	MixedRate int

	Station string

	Error error
}

type Event struct {
	Name       string `json:"name"`
	DateString string `json:"date"`

	Date time.Time

	LeftDays int
}
