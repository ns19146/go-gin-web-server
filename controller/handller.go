package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/render-examples/go-gin-web-server/model"
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
	db.SingularTable(true)
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

func InsertPlayerInf(c *gin.Context) {
	var player model.PlayerInf
	player.PlayerName = c.PostForm("player_name")
	player.Gender = c.PostForm("gender")
	player.BirthDate = c.PostForm("birth_date")
	player.Team = c.PostForm("team")
	player.Number = c.PostForm("number")
	db := dbInit()
	db.Save(&player)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ShowPlayerList(c *gin.Context) {
	var players []model.PlayerInf
	db := dbInit()
	db.Find(&players)
	c.HTML(http.StatusOK, "player_list.html", gin.H{
		"players": players,
	})
}

func SearchPlayerInf(c *gin.Context) {
	var players []model.PlayerInf
	name := c.PostForm("player_name")
	db := dbInit()
	db.Where("player_name=?", name).Find(&players)
	c.HTML(http.StatusOK, "player_list.html", gin.H{
		"players": players,
		"name":    name,
	})
}

func InsertGameInf(c *gin.Context) {
	var game model.GameInf
	game.GameName = c.PostForm("game_name")
	game.Date = c.PostForm("date")
	game.Location = c.PostForm("location")
	db := dbInit()
	db.Save(&game)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ShowGameinfForm(c *gin.Context) {
	var games []model.GameInf
	db := dbInit()
	db.Find(&games)
	c.HTML(http.StatusOK, "player_inf.html", gin.H{
		"games": games,
	})
}

func EditPlayerInf(c *gin.Context) {
	var player model.PlayerInf
	id, _ := strconv.Atoi(c.PostForm("id"))
	db := dbInit()
	db.Where("player_id = ?", id).First(&player)
	c.HTML(http.StatusOK, "edit_player_inf.html", gin.H{
		"player": player,
	})
}

func UpdatePlayerInf(c *gin.Context) {
	var player model.PlayerInf
	id, _ := strconv.Atoi(c.PostForm("id"))
	player.PlayerID = id
	player.PlayerName = c.PostForm("name")
	player.BirthDate = c.PostForm("birthdate")
	player.Gender = c.PostForm("gender")
	player.Team = c.PostForm("team")
	player.Number = c.PostForm("number")
	db := dbInit()
	db.Where("player_id = ?", id).Save(&player)
	c.Redirect(http.StatusMovedPermanently, "/playerinf/search")
}

func ShowGameList(c *gin.Context) {
	var games []model.GameInf
	db := dbInit()
	db.Find(&games)
	c.HTML(http.StatusOK, "game_list.html", gin.H{
		"games": games,
	})
}

func SearchGameInf(c *gin.Context) {
	var games []model.GameInf
	name := c.PostForm("game_name")
	db := dbInit()
	db.Where("game_name=?", name).Find(&games)
	c.HTML(http.StatusOK, "game_list.html", gin.H{
		"games": games,
		"name":  name,
	})
}

func EditGameInf(c *gin.Context) {
	var game model.GameInf
	id, _ := strconv.Atoi(c.PostForm("id"))
	db := dbInit()
	db.Where("game_id = ?", id).First(&game)
	c.HTML(http.StatusOK, "edit_game_inf.html", gin.H{
		"game": game,
	})
}

func UpdateGameInf(c *gin.Context) {
	var game model.GameInf
	id, _ := strconv.Atoi(c.PostForm("id"))
	game.GameName = c.PostForm("name")
	game.Date = c.PostForm("date")
	game.Location = c.PostForm("location")
	db := dbInit()
	db.Where("game_id = ?", id).Save(&game)
	c.Redirect(http.StatusMovedPermanently, "/gameinf/search")
}
