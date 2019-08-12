package model

import (
	"fmt"

	"github.com/go-pg/pg"
)

func (r *Report) Create(db *pg.DB) (err error) {
	err = db.Insert(r)
	return
}

func (r *Report) String() string {
	var format string
	if r.ReportType == "위반" {
		format = "대상 학생: %d학년 %d반 %s\\n%v에 %s(%s)을(를) 하였습니다."
		return fmt.Sprintf(format, r.TargetStudent.Grade, r.TargetStudent.Class,
			r.TargetStudent.Name, r.When, r.What, r.Detail)
	} else {
		format = "대상 학생: %d학년 %d반 %s\\n%v에 %s을(를) 하였습니다."
		return fmt.Sprintf(format, r.TargetStudent.Grade, r.TargetStudent.Class,
			r.TargetStudent.Name, r.When, r.Detail)
	}
}

func FindReportByDetailAndUserID(db *pg.DB, userID, detail string) (r Report, err error) {
	err = db.Model(&r).Where("user_id = ?", userID).Where("detail = ?", detail).Limit(1).Select()
	return
}

func (r *Report) Cancel(db *pg.DB) (err error) {
	_, err = db.Model(r).Column("was_canceled").WherePK().Update()
	return
}
