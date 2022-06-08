package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello go gin web.",
	})
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)

	var router = gin.Default()
	var port = ":8080"

	router.GET("/", index) // handler for root

	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uid":           c.PostForm("uid"),
			"pwd":           c.PostForm("pwd"),
			"rememberMe":    c.DefaultPostForm("rememberMe", "0"),
			"Authorization": c.GetHeader("Authorization"),
		})
	})

	router.Run(port)
}
