package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartWebServer() *gin.Engine {
	r := gin.Default()

	//リソース読み込み
	r.Static("/css", "/assets/css")
	r.LoadHTMLGlob("templates/*.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	//index
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "Hello World!",
		})
	})

	r.GET("/list", ShowTables)
	//r.GET("/upload", UploadCsv)
	r.GET("/migration", Migration)
	//r.POST("/upload", OpenCsv)

	//選手情報関連
	player := r.Group("/playerinf")
	{
		player.GET("/insert", ShowPlayerinfForm)
		player.POST("/insert", InsertPlayerInf)
		player.GET("/search", ShowPlayerList)
		player.POST("/search", SearchPlayerInf)
		player.POST("/edit", EditPlayerInf)
		player.POST("/update", UpdatePlayerInf)
		player.POST("/delete", DeletePlayerInf)
	}

	//大会情報関連
	game := r.Group("/gameinf")
	{
		game.GET("insert", func(c *gin.Context) {
			c.HTML(http.StatusOK, "game_inf.html", nil)
		})
		game.GET("/search", ShowGameList)
		game.POST("/insert", InsertGameInf)
		game.POST("/search", SearchGameInf)
		game.POST("/edit", EditGameInf)
		game.POST("/update", UpdateGameInf)
		game.POST("/delete", DeleteGameInf)
	}

	//スコア関連
	score := r.Group("/score")
	{
		score.GET("/insert", ShowScoreForm)
		score.POST("/insert", InsertScore)
		score.POST("/search", ShowScoreList)
		score.POST("/edit", nil)
		score.POST("/update", nil)
	}

	csv := r.Group("/csv")
	{
		csv.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload.html", nil)
		})
		csv.POST("/upload", nil)
	}

	return r
}
