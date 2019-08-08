package model

import "github.com/go-pg/pg"

func (f *Feedback) Create(db *pg.DB) (err error) {
	err = db.Insert(f)
	return
}
