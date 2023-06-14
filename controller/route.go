package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartWebServer() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/list", ShowTables)
	//r.GET("/upload", UploadCsv)
	r.GET("/migration", Migration)
	//r.POST("/upload", OpenCsv)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "Hello World!",
		})
	})
	r.GET("/playerinf/insert", ShowGameinfForm)
	r.POST("/playerinf/insert", InsertPlayerInf)
	r.GET("/playerinf/search", ShowPlayerList)
	r.POST("/playerinf/search", SearchPlayerInf)
	r.POST("playerinf/edit", EditPlayerInf)
	r.POST("/playerinf/update", UpdatePlayerInf)
	r.GET("/gameinf/insert", func(c *gin.Context) {
		c.HTML(http.StatusOK, "game_inf.html", nil)
	})
	r.POST("/gameinf/insert", InsertGameInf)
	return r
}
