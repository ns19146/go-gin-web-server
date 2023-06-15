package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/model"
	"net/http"
	"strconv"
)

func ShowGameList(c *gin.Context) {
	var games []model.GameInf
	db := dbInit()
	db.Find(&games)
	c.HTML(http.StatusOK, "game_list.html", gin.H{
		"games": games,
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

func SearchGameInf(c *gin.Context) {
	var games []model.GameInf
	name := c.PostForm("game_name")
	db := dbInit()
	db.Where("game_name = ?", name).Find(&games)
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
	game.GameID = id
	game.GameName = c.PostForm("name")
	game.Date = c.PostForm("date")
	game.Location = c.PostForm("location")
	db := dbInit()
	db.Where("game_id = ?", id).Save(&game)
	c.Redirect(http.StatusMovedPermanently, "/gameinf/search")
}

func DeleteGameInf(c *gin.Context) {
	var game model.GameInf
	id, _ := strconv.Atoi(c.PostForm("id"))
	db := dbInit()
	db.Where("game_id = ?", id).Delete(&game)
	c.Redirect(http.StatusMovedPermanently, "/gameinf/search")
}
