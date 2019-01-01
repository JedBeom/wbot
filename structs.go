package main

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
