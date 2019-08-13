package model

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
)

func (r *Report) Create(db *pg.DB) (err error) {
	err = db.Insert(r)
	return
}

func (r *Report) String() string {
	format := "[정보]\\n대상 학생: %d학년 %d반 %s\\n시각: %v\\n"
	if r.ReportType == "위반" {
		format += "상세: %s(%s)"
		return fmt.Sprintf(format, r.TargetStudent.Grade, r.TargetStudent.Class,
			r.TargetStudent.Name, timeToStr(r.When), r.What, r.Detail)
	} else {
		format += "상세: %s"
		return fmt.Sprintf(format, r.TargetStudent.Grade, r.TargetStudent.Class,
			r.TargetStudent.Name, timeToStr(r.When), r.Detail)
	}
}

func timeToStr(t time.Time) string {
	return t.Format("2006/01/02 15:04")
}

func FindReportByDetailAndUserID(db *pg.DB, userID, detail string) (r Report, err error) {
	err = db.Model(&r).Where("user_id = ?", userID).Where("detail = ?", detail).Limit(1).Select()
	return
}

func (r *Report) Cancel(db *pg.DB) (err error) {
	_, err = db.Model(r).Column("was_canceled").WherePK().Update()
	return
}
