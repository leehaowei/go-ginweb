package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"strings"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello go gin web.",
	})
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,password"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required,oneof=ADMIN USER"`
}

func verifyPassword(field validator.FieldLevel) bool {
	var regex = regexp.MustCompile("\\w{8,}")
	var password = field.Field().String()
	if regex.MatchString(password) {
		if strings.Index(password, "!") > -1 {
			return true
		}
	}
	return false
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)

	var router = gin.Default()
	var port = ":8080"

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", verifyPassword)
	}

	router.GET("/", index) // handler for root

	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name":     user.Name,
			"password": user.Password,
			"email":    user.Email,
			"role":     user.Role,
		})
	})

	router.Run(port)
}
