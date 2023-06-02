package model

import "time"

type Score struct {
	Gamename  string `gorm:"type:varchar(30); primary_key" json:"gamename"`
	Number    string `gorm:"type:varchar(10); primary_key" json:"number"`
	Dist      string `gorm:"type:varchar(10); primary_key" json:"dist"`
	Xring     int    `json:"xring"`
	Set       int    `gorm:"primary_key" json:"set"`
	SetScore  int    `json:"setscore"`
	Distscore int    `json:"distscore"`
	Total     int    `json:"total"`
}

type Gameinf struct {
	Gamename string    `gorm:"type:varchar(30); primary_key" json:"gamename"`
	Date     time.Time `json:"date"`
	Locate   string    `gorm:"type:varchar(20)" json:"locate"`
}

type Playerinf struct {
	Name     string `gorm:"type:varchar(20)" json:"name"`
	Age      int    `json:"age"`
	Gen      string `gorm:"type:varchar(10)" json:"gen"`
	Team     string `gorm:"type:varchar(20)" json:"team"`
	Number   string `gorm:"type:varchar(10); primary_key" json:"number"`
	Gamename string `gorm:"type:varchar(30); primary_key" json:"gamename"`
}

type TestModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Gen  string `json:"gen"`
	Team string `json:"team"`
}
