package api

import (
	"seed/models"
	"seed/modules"

	"github.com/gin-gonic/gin"
)

type schemes int

var Schemes schemes

func (schemes) Save(c *gin.Context) {
	var args struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	if c.Bind(&args) != nil {
		return
	}
	s := new(models.Schemes)
	s.ID = args.ID
	s.Name = args.Name
	if err := s.Save(); err != nil {
		JSON(c, err)
	} else {
		JSON(c, "success")
	}
}

func (schemes) List(c *gin.Context) {
	list, err := models.SchemesDefault.List()
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, list)
	}
}

func (schemes) GetByID(c *gin.Context) {
	var args struct {
		ID int `form:"id"`
	}
	if c.Bind(&args) != nil {
		return
	}
	data, err := modules.Schemes.GetDetail(args.ID)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (schemes) SaveDetail(c *gin.Context) {
	var args struct {
		SID             int                   `json:"sid"`
		FixedCostList   []*models.FixedCost   `json:"fixed_cost_list"`
		CoursePriceList []*models.CoursePrice `json:"course_price_list"`
		StaffList       []*models.Staff       `json:"staff_list"`
	}
	if c.Bind(&args) != nil {
		return
	}
	err := modules.Schemes.SaveDetail(args.SID, args.FixedCostList, args.CoursePriceList, args.StaffList)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, "success")
	}
}
