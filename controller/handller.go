package controller

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/render-examples/go-gin-web-server/model"
	"io"
	"log"
	"net/http"
	"strconv"
)

func dbInit() *gorm.DB {
	dsn := "host=dpg-chdin0bhp8u3v70u25og-a.singapore-postgres.render.com port=5432 user=nittc2023_j5exp_g2 dbname=nittc2023_j5exp_g2 password=uEeRkwJRnQufgEbbF3EFnrJUm0BDJRzP"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	db.SingularTable(true)

	db.AutoMigrate(
		&model.TestModel{},
	)
	return db
}

func CreateTable(_ *gin.Context) {
	var game model.Gameinf
	var player model.Playerinf
	var score model.Score
	db := dbInit()
	db.CreateTable(&game)
	db.CreateTable(&player)
	db.Model(&player).AddForeignKey("gamename", "gameinf(gamename)", "RESTRICT", "RESTRICT")
	db.CreateTable(&score)
	db.Model(&score).AddForeignKey("gamename", "gameinf(gamename)", "RESTRICT", "RESTRICT")
	db.Model(&score).AddForeignKey("number", "playerinf(number)", "RESTRICT", "RESTRICT")
}

func ShowTables(c *gin.Context) {
	var game []model.Gameinf
	var player []model.Playerinf
	var score []model.Score
	db := dbInit()
	db.Find(&game)
	fmt.Println(game[0].Gamename)
	db.Find(&player)
	db.Find(&score)
	c.HTML(http.StatusOK, "show.html", gin.H{
		"game":   game,
		"player": player,
		"score":  score,
	})
}

func InsertTable(c *gin.Context) {
	var model model.Score
	db := dbInit()
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	db.Create(&model)
}

func UploadCsv(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func OpenCsv(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	var line []string
	for {
		var model model.TestModel
		line, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		id, _ := strconv.Atoi(line[0])
		if id != 0 {
			log.Println(line)
			model.Id = id
			model.Name = line[1]
			model.Gen = line[2]
			model.Team = line[3]
		}
		db := dbInit()
		db.Save(&model)
	}
	defer c.Redirect(http.StatusSeeOther, "https://nittc2023-j5exp-g2-2pkv.onrender.com/")
}

/*
	ReadAllを用いる場合
	for _, fields := range line {
		fmt.Println(fields)
	}
*/
