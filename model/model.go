package model

type PlayerInf struct {
	PlayerID   int    `gorm:"primary_key" json:"playerid"`
	PlayerName string `gorm:"varchar(10)" json:"gamename"`
	BirthDate  string `gorm:"varchar(20)" json:"birthdate"`
	Gender     string `gorm:"varchar(10)" json:"gender"`
	Team       string `gorm:"varchar(20)" json:"team"`
	Number     string `gorm:"varchar(5)" json:"number"`
}

type GameInf struct {
	GameID   int    `gorm:"primary_key" json:"id"`
	GameName string `gorm:"varchar(50)" json:"gamename"`
	Date     string `json:"date"`
	Location string `gorm:"varchar(30)" json:"location"`
}

type EntryPlayerInf struct {
	GameID   int `gorm:"primary_key" json:"gameid"`
	PlayerID int `gorm:"primary_key" json:"playerid"`
	ScoreID  int `json:"scoreid"`
}

type Score struct {
	ScoreID   int    `gorm:"primary_key; unique" json:"scoreid"`
	Distance  string `gorm:"varchar(5); primary_key" json:"distance"`
	SetNumber int    `gorm:"primary_key" json:"setnumber"`
	XRIng     int    `json:"xring"`
	Shot1     int    `json:"shot1"`
	Shot2     int    `json:"shot2"`
	Shot3     int    `json:"shot3"`
	Shot4     int    `json:"shot4"`
	Shot5     int    `json:"shot5"`
	Shot6     int    `json:"shot6"`
}

type TestModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Gen  string `json:"gen"`
	Team string `json:"team"`
}
