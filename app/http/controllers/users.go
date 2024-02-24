package controllers

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/exceptions"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	user.LastLoginTime = time.Now()
	user.IconPath = ""
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
	iconpath := user.IconPath
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	user.IconPath = iconpath
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

func UploadUserIcon(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := user.FindUserByID(id); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err := utils.DeleteFile(user.IconPath)
	if err != nil && user.IconPath != "" {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	var image utils.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	outputFile := "storage/users/" + id
	savePath, err := utils.SaveDecodedImage(image.EncodedData, outputFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	user.IconPath = savePath
	if err := user.UpdateUserByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func DeleteUserIcon(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := user.FindUserByID(id); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	
	if err := utils.DeleteFile(user.IconPath); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	user.IconPath = ""
	if err := user.UpdateUserByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}