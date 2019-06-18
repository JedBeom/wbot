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
	Weekday    string
	NormalText string
}

// 전역 변수로 쓰이는 미세먼지 구조체
type HangulQ struct {
	Pm10 string // 미세먼지 등급
	Pm25 string // 초미세먼지 등급

	MixedRate  int // 둘중에 더 안좋은 등급 저장
	TimeString string

	Station string // 측정소 이름

	Error error // 미세먼지를 가져오는 중의 에러
}

// 학사 일정
type Event struct {
	// events.json에서 가져옴
	Name       string `json:"name"` // 이름
	DateString string `json:"date"` // yyyy-mm-dd
	After      int    `json:"after,omitempty"`

	// 위 필드를 가공해 얻음
	Date     time.Time // DateString 에서 파싱된 Go Time 구조체
	MMDD     string
	LeftDays int  // 남은 날 수
	IsDDAY   bool // 오늘이 혁명의 그 날입니까
}

// 응답 구조체들

type Button struct {
	Action string `json:"action"`
	Label  string `json:"label"`
	URL    string `json:"webLinkUrl"`
}

type Thumbnail struct {
	ImgURL string `json:"imageUrl"`
}

type BasicCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	Profile *Profile `json:"profile,omitempty"`

	Social *Social `json:"social,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
	Buttons   []Button   `json:"buttons"`
}

type Profile struct {
	ImgURL   string `json:"imageUrl"`
	Nickname string `json:"nickname"`
}

type Social struct {
	Like    int `json:"like"`
	Comment int `json:"comment"`
	Share   int `json:"share"`
}
