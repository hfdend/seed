package models

import "seed/cli"

type CoursePrice struct {
	Model
	SID          int     `json:"sid" gorm:"column:sid;index"`
	Name         string  `json:"name"`
	OPrice       float64 `json:"o_price"`
	Price        float64 `json:"price"`
	Numbers      int     `json:"numbers"`
	Student      int     `json:"student"`
	PerNumber    int     `json:"per_number"`
	CourseAmount float64 `json:"course_amount"`
	Percentage   float64 `json:"percentage"`
}

var CoursePriceDefault CoursePrice

func (CoursePrice) TableName() string {
	return "course_price"
}

func (f CoursePrice) DelBySID(sid int) error {
	return cli.DB.Delete(CoursePrice{}, "sid = ?", sid).Error
}

func (f *CoursePrice) Save() error {
	return cli.DB.Save(f).Error
}

func (f CoursePrice) ListBySID(sid int) (list []*CoursePrice, err error) {
	err = cli.DB.Where("sid = ?", sid).Find(&list).Error
	return
}
