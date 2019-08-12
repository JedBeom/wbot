package model

import "github.com/go-pg/pg"

func (s *Student) Find(db *pg.DB) (err error) {
	err = db.Model(s).WhereStruct(s).Select()
	return
}

func (s *Student) GetByID(db *pg.DB) (err error) {
	err = db.Model(s).Where("id = ?", s.ID).Select()
	return
}
