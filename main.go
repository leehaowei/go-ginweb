package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	var router = gin.Default()
	var port = ":8080"

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Token")
		if len(token) == 0 {
			c.Writer.Header().Add("Token", "abcde.12345.XYZ")
		} else {
			c.Writer.Header().Add("Token", token+"-ABC")
		}
		fmt.Println("> access logging...")
	})

	router.StaticFile("/", "./static/index.html")

	router.Run(port)
}
