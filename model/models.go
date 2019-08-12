package model

import "time"

type Report struct {
	ID int

	// 건의, 도움필요, 위반, 칭찬
	ReportType string

	// 유저 입력
	TargetStudentID int // 누가
	TargetStudent   *Student
	What            string    // 무엇을
	When            time.Time // 언제
	Detail          string    // 자세히

	HistoryID int
	History   *History

	UserID string
	User   *User

	WasCanceled bool

	CreatedAt time.Time `sql:"default:now()"`
}

type Student struct {
	ID int

	Grade  int    `sql:",unique:identify"`
	Class  int    `sql:",unique:identify"`
	Number int    `sql:",unique:identify"`
	Name   string `sql:",unique:identify"`

	CardID string `sql:",unique"`
}

type User struct {
	ID string `sql:",pk"`

	StudentID int
	Student   Student

	CreatedAt time.Time `sql:"default:now()"`
}

type Feedback struct {
	ID int

	HistoryID int
	History   *History

	UserID string
	User   *User

	Text string

	CreatedAt time.Time `sql:"default:now()"`
}

// Logging
type History struct {
	ID int

	User   *User
	UserID string

	Utterance string
	BlockID   string
	BlockName string

	Params map[string]string

	Date time.Time `sql:"default:now()"`

	ContextDetail string `sql:"-"`
}
