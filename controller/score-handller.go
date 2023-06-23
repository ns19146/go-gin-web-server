package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/model"
	"log"
	"net/http"
	"strconv"
)

func ShowScoreList(c *gin.Context) {
	var entry []model.EntryPlayerInf
	var game model.GameInf
	var details []model.ScoreDetail
	var game_id int
	game_id, _ = strconv.Atoi(c.PostForm("id"))
	db := dbInit()
	db.Where("game_id = ?", game_id).Find(&entry)
	db.Where("game_id = ?", game_id).First(&game)

	for _, res := range entry {
		var player model.PlayerInf
		var detail model.ScoreDetail
		var score model.Score

		player_id := res.PlayerID
		log.Print("player_id = ", +player_id)
		db.Where("player_id = ?", player_id).Find(&player)
		score_id := res.ScoreID
		db.Where("score_id = ?", score_id).Find(&score)
		log.Print("player_id = ", +player.PlayerID, "game_id = ", +game.GameID, "score_id = ", +score.ScoreID)
		detail.GameID = game.GameID
		detail.GameName = game.GameName
		detail.PlayerID = player_id
		detail.PlayerName = player.PlayerName
		detail.ScoreID = score.ScoreID
		detail.Distance = score.Distance
		detail.SetNumber = score.SetNumber
		detail.XRing = score.XRIng
		detail.Shot1 = score.Shot1
		detail.Shot2 = score.Shot2
		detail.Shot3 = score.Shot3
		detail.Shot4 = score.Shot4
		detail.Shot5 = score.Shot5
		detail.Shot6 = score.Shot6
		details = append(details, detail)
	}
	c.HTML(http.StatusOK, "score_list.html", gin.H{
		"details": details,
	})
}
