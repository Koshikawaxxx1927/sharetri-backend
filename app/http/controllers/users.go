package controllers

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/exceptions"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	user.LastLoginTime = time.Now()

	if err := user.CreateUser(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func FindUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := user.FindUserByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := user.FindUserByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if err := user.UpdateUserByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func DeleteUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := user.FindUserByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := user.DeleteUserByID(id); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}