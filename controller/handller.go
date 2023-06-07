package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/render-examples/go-gin-web-server/model"
	"net/http"
)

func dbInit() *gorm.DB {
	dsn := "host=dpg-chdin0bhp8u3v70u25og-a.singapore-postgres.render.com port=5432 user=nittc2023_j5exp_g2 dbname=nittc2023_j5exp_g2 password=uEeRkwJRnQufgEbbF3EFnrJUm0BDJRzP"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migration(_ *gin.Context) {
	var player model.PlayerInf
	var game model.GameInf
	var score model.Score
	var entry model.EntryPlayerInf

	db := dbInit()
	db.SingularTable(true)
	db.CreateTable(&player)
	db.CreateTable(&game)
	db.CreateTable(&score)
	db.CreateTable(&entry)
	db.Model(&entry).AddForeignKey("game_id", "game_inf(game_id)", "RESTRICT", "RESTRICT")
	db.Model(&entry).AddForeignKey("player_id", "player_inf(player_id)", "RESTRICT", "RESTRICT")
	db.Model(&entry).AddForeignKey("score_id", "score(score_id)", "RESTRICT", "RESTRICT")
}

func ShowTables(c *gin.Context) {
	var game []model.GameInf
	var player []model.PlayerInf
	var entry []model.EntryPlayerInf
	var score []model.Score
	db := dbInit()
	db.Find(&game)
	db.Find(&player)
	db.Find(&score)
	db.Find(&entry)
	c.HTML(http.StatusOK, "show.html", gin.H{
		"game":   game,
		"player": player,
		"score":  score,
		"entry":  entry,
	})
}

func UploadCsv(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

/*
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
