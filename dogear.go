package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.html", gin.H{
			"title": "New Bookmark",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
