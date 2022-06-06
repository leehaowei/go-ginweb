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

	router.GET("/", index)

	router.Run(port)
}
