package main

import "time"

// 카카오에서 보내오는 정보
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

// 전역 변수로 쓰이는 미세먼지 구조체
type HangulQ struct {
	Pm10 string // 미세먼지 등급
	Pm25 string // 초미세먼지 등급

	MixedRate int // 둘중에 더 안좋은 등급 저장

	Station string // 측정소 이름

	Error error // 미세먼지를 가져오는 중의 에러
}

// 학사 일정
type Event struct {
	// events.json에서 가져옴
	Name       string `json:"name"` // 이름
	DateString string `json:"date"` // yyyy-mm-dd

	// 위의 두 필드를 가공해 얻음
	Date time.Time // DateString에서 파싱된 Go Time 구조체

	LeftDays int // 남은 날 수
}
