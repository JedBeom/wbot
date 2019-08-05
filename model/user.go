package model

import (
	"github.com/go-pg/pg"
)

func (u *User) Create(db *pg.DB) error {
	err := db.Insert(u)
	return err
}

func GetUserByID(db *pg.DB, id string) (u User, err error) {
	err = db.Model(&u).Where("id = ?", id).Select()
	return
}
