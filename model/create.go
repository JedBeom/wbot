package model

import (
	"github.com/go-pg/pg"
)

func CreateTables(db *pg.DB) error {
	for _, model := range []interface{}{&Student{}, &User{}, &History{}, &Report{}, &Feedback{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
