package models

import (
	"fmt"
	"seed/cli"
	"testing"
)

func TestSchemes_Save(t *testing.T) {
	cli.Init()
	s := &Schemes{
		Name: "哈哈",
	}
	err := s.Save()
	fmt.Println(err)
	fmt.Println(s)
}

func TestSchemes_List(t *testing.T) {
	cli.Init()
	list, err := (Schemes{}).List()
	fmt.Println(err)
	fmt.Println(list)
}
