package models

import (
	"seed/cli"
)

// 固定成本
type FixedCost struct {
	Model
	SID    int     `json:"sid" gorm:"column:sid;index"`
	Name   string  `json:"name"`
	Day    int     `json:"day"`
	Amount float64 `json:"amount"`
}

var FixedCostDefault FixedCost

func (FixedCost) TableName() string {
	return "fixed_costs"
}

func (f FixedCost) DelBySID(sid int) error {
	return cli.DB.Delete(FixedCost{}, "sid = ?", sid).Error
}

func (f *FixedCost) Save() error {
	return cli.DB.Save(f).Error
}

func (f FixedCost) ListBySID(sid int) (list []*FixedCost, err error) {
	err = cli.DB.Where("sid = ?", sid).Find(&list).Error
	return
}
