package controllers

import (
	"log"
	"time"
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatal("Faild to bind json")
	}
	user.LastLoginTime = time.Now()

	if err := user.CreateUser(); err != nil {
		log.Fatal("Failed to create new user")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": user,
	})
}

func FindUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	user.FindUserByID(id)

	c.JSON(http.StatusOK, gin.H{
		"code": user,
	})
}

func UpdateUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	user.FindUserByID(id)

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatal("Failed to bind json")
	}
	if err := user.UpdateUserByID(); err != nil {
		log.Fatal("Failed to update user")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": user,
	})
}

func DeleteUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := user.DeleteUserByID(id); err != nil {
		log.Fatal("Failed to delete the user")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": user,
	})
}