package models

import "seed/cli"

type Staff struct {
	Model
	SID    int     `json:"sid" gorm:"column:sid;index"`
	Name   string  `json:"name"`
	Number int     `json:"number"`
	Duty   string  `json:"duty"`
	Info   string  `json:"info"`
	Wages  float64 `json:"wages"`
}

var StaffDefault Staff

func (Staff) TableName() string {
	return "staffs"
}

func (f Staff) DelBySID(sid int) error {
	return cli.DB.Delete(Staff{}, "sid = ?", sid).Error
}

func (f *Staff) Save() error {
	return cli.DB.Save(f).Error
}

func (f Staff) ListBySID(sid int) (list []*Staff, err error) {
	err = cli.DB.Where("sid = ?", sid).Find(&list).Error
	return
}
