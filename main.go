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

type uriNewsBinding struct {
	Region   string `uri:"region" binding:"required"`
	Category string `uri:"category" binding:"required"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)

	var router = gin.Default()
	var port = ":8080"

	router.GET("/", index) // handler for root

	router.GET("news/:region/:category", func(c *gin.Context) {
		var binding uriNewsBinding

		if err := c.ShouldBindUri(&binding); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"region":   binding.Region,
			"category": binding.Category,
		})
	})

	router.Run(port)
}
