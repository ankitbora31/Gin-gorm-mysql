package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rest/API/controllers"
	"rest/API/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "first APi"})
	})

	models.InitDb()

	r.GET("/users", controllers.GetUsers)
	r.POST("/createUser", controllers.CreateUser)

	r.Run()
}
