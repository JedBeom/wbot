package model

import "time"

type Report struct {
	ID int

	// notify, praise, violation, assistant, proposal
	Type string

	UserID string
	User   *User

	HistoryID int
	History   *History

	TargetStudent   *Student
	TargetStudentID int

	IsRead bool

	CreatedAt time.Time `sql:"default:now()"`
}

type Student struct {
	ID int

	Grade  int    `sql:",unique:identify"`
	Class  int    `sql:",unique:identify"`
	Number int    `sql:",unique:identify"`
	Name   string `sql:",unique:identify"`

	CardID string `sql:",unique"`

	CreatedAt time.Time `sql:"default:now()"`
}

type User struct {
	ID        string `sql:",pk" sql:",unique"`
	Student   *Student
	StudentID string

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
}
