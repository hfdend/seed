package modules

import (
	"seed/models"
)

type schemes int

var Schemes schemes

func (schemes) SaveDetail(sid int, fixedCostList []*models.FixedCost, coursePriceList []*models.CoursePrice, staffs []*models.Staff) (err error) {
	if err = models.FixedCostDefault.DelBySID(sid); err != nil {
		return
	}
	if err = models.CoursePriceDefault.DelBySID(sid); err != nil {
		return
	}
	if err = models.StaffDefault.DelBySID(sid); err != nil {
		return
	}
	for _, v := range fixedCostList {
		v.SID = sid
		if err = v.Save(); err != nil {
			return
		}
	}
	for _, v := range coursePriceList {
		v.SID = sid
		if err = v.Save(); err != nil {
			return
		}
	}
	for _, v := range staffs {
		v.SID = sid
		if err = v.Save(); err != nil {
			return
		}
	}
	return
}

func (schemes) GetDetail(sid int) (data *models.Schemes, err error) {
	if data, err = models.SchemesDefault.GetByID(sid); err != nil {
		return
	}
	if data.FixedCosts, err = models.FixedCostDefault.ListBySID(sid); err != nil {
		return
	}
	if data.CoursePrice, err = models.CoursePriceDefault.ListBySID(sid); err != nil {
		return
	}
	if data.Staffs, err = models.StaffDefault.ListBySID(sid); err != nil {
		return
	}
	return
}
