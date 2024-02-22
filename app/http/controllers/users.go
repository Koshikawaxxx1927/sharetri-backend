package controllers

import (
	"log"
	"time"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	iconpath := c.PostForm("iconpath")

	fmt.Println(name)
	fmt.Println(iconpath)

	user := models.User{
		Name: name,
		IconPath: iconpath,
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

func UpdateUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	user.FindUserByID(id)

	name := c.PostForm("name")
	iconpath := c.PostForm("iconpath")
	user.Name = name
	user.IconPath = iconpath
	user.UpdateUserByID()

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