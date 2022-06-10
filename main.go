package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	var router = gin.Default()
	var port = ":8080"

	router.LoadHTMLGlob("./templates/*")

	router.GET("/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"name": "frankfurt",
			"role": "ADMIN",
			"age":  25,
		})
	})

	router.StaticFile("/", "./static/index.html")

	router.Run(port)
}
