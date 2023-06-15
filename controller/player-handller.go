package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/model"
	"net/http"
	"strconv"
)

func ShowPlayerinfForm(c *gin.Context) {
	var games []model.GameInf
	db := dbInit()
	db.Find(&games)
	c.HTML(http.StatusOK, "player_inf.html", gin.H{
		"games": games,
	})
}

func InsertPlayerInf(c *gin.Context) {
	var player model.PlayerInf
	player.PlayerName = c.PostForm("player_name")
	player.Gender = c.PostForm("gender")
	player.BirthDate = c.PostForm("birth_date")
	player.Team = c.PostForm("team")
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
	db := dbInit()
	db.Where("player_id = ?", id).Save(&player)
	c.Redirect(http.StatusMovedPermanently, "/playerinf/search")
}

func DeletePlayerInf(c *gin.Context) {
	var player model.PlayerInf
	id, _ := strconv.Atoi(c.PostForm("id"))
	db := dbInit()
	db.Where("player_id = ?", id).Delete(&player)
	c.Redirect(http.StatusMovedPermanently, "/playerinf/search")
}
