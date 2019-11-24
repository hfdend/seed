package models

import (
	"fmt"
	"seed/cli"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	SureYes Sure = 1
	SureNo  Sure = -1
	SureNil Sure = 0
)

type Sure int

func (s Sure) Bool() bool {
	return s == SureYes
}

func Begin() *gorm.DB {
	return cli.DB.Begin()
}

type Model struct {
	ID int `json:"id" gorm:"primary_key"`
}

func DBInsertIgnore(dbq *gorm.DB, obj interface{}) (int64, error) {
	scope := dbq.NewScope(obj)
	fields := scope.Fields()
	quoted := make([]string, 0, len(fields))
	placeholders := make([]string, 0, len(fields))
	for i := range fields {
		if fields[i].IsPrimaryKey && fields[i].IsBlank {
			continue
		}
		quoted = append(quoted, scope.Quote(fields[i].DBName))
		placeholders = append(placeholders, scope.AddToVars(fields[i].Field.Interface()))
	}

	scope.Raw(fmt.Sprintf("INSERT IGNORE INTO %s (%s) VALUES (%s)",
		scope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholders, ", ")))

	result, err := scope.SQLDB().Exec(scope.SQL, scope.SQLVars...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func AutoMigrate() {
	tables := []interface{}{
		&Schemes{},
		&FixedCost{},
		&CoursePrice{},
		&Staff{},
	}
	for _, v := range tables {
		cli.DB.AutoMigrate(v)
	}
}
