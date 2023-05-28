package models

import (
	"github.com/jinzhu/gorm"
)

type TestModel struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Score struct {
	Onescore  int
	Distscore int
	Total     int
	Athname   string `gorm:"type:varchar20;primary_key"`
}
