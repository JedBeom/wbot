package model

import "github.com/go-pg/pg"

func (h *History) Create(db *pg.DB) (err error) {
	err = db.Insert(h)
	return
}
