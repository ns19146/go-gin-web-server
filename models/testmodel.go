package models

import (
	"github.com/jinzhu/gorm"
)

type TestModel struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}
