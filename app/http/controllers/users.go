package controllers

import (
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateUser(c *gin.Context) {
	user := models.User{
		Name: "千葉",
		IconPath: "./resources/usericon/me.jpeg",
		LastLoginTime: time.Now(),
	}
	
	if err := user.CreateUser(); err != nil {
		log.Fatal("Failed to create new user")
	}
	
	// config.GetDB().First(&user, 3)

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