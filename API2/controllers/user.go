package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rest/API/models"
)

type CreateUserInput struct {
	Name    string `json:"name" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	models.Db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Gender: input.Gender, Age: input.Age, Address: input.Address}
	models.Db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
