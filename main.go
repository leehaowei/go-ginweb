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

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"userId": c.Query("id"),
			"sex":    c.Query("sex"),
			"age":    c.DefaultQuery("age", "18"),
		})
	})

	//router.GET("/users/:id", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"userId": c.Param("id"),
	//	})
	//})
	//router.GET("news/:id", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"newsID": c.Param("id"),
	//	})
	//})

	// define router Group - V1 (version1)
	//var apiV1 = router.Group("/v1")
	//apiV1.GET("/user", func(c *gin.Context) {
	//	c.String(http.StatusOK, "User GET service v1")
	//})
	//apiV1.GET("news", func(c *gin.Context) {
	//	c.String(http.StatusOK, "NEWS GET service v1")
	//})

	// define router Group - V2 (version2)
	//var apiV2 = router.Group("/v2")
	//apiV2.GET("/user", func(c *gin.Context) {
	//	c.String(http.StatusOK, "User GET service v2")
	//})
	//apiV2.GET("/news", func(c *gin.Context) {
	//	c.String(http.StatusOK, "NEWS GET service v2")
	//})

	router.Run(port)
}
