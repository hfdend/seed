package models

import (
	"seed/cli"
	"time"
)

type Schemes struct {
	Model
	Name    string    `json:"name"`
	Created time.Time `json:"created"`

	FixedCosts  []*FixedCost   `json:"fixed_costs,omitempty"`
	CoursePrice []*CoursePrice `json:"course_price,omitempty"`
	Staffs      []*Staff       `json:"staffs,omitempty"`
}

var SchemesDefault Schemes

func (Schemes) TableName() string {
	return "schemes"
}

func (s *Schemes) Save() error {
	s.Created = time.Now()
	return cli.DB.Create(s).Error
}

func (Schemes) List() (list []*Schemes, err error) {
	err = cli.DB.Find(&list).Error
	return
}

func (Schemes) GetByID(id int) (data *Schemes, err error) {
	data = new(Schemes)
	err = cli.DB.Where("id = ?", id).Find(data).Error
	return
}
