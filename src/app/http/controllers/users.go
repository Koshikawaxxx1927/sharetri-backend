package controllers

import (
	// "time"
	"io"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/exceptions"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/utils"
	// "strconv"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// user.LastLoginTime = time.Now()
	user.Iconpath = ""
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
	userid := c.Param("userid")
	if err := user.FindUserByID(userid); err == exceptions.NotFound {
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
	userid := c.Param("userid")
	if err := user.FindUserByID(userid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	iconpath := user.Iconpath
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	user.Iconpath = iconpath
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
	userid := c.Param("userid")

	if err := user.FindUserByID(userid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := user.DeleteUserByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func UploadUserIcon(c *gin.Context) {
	var user models.User
	userid := c.Param("userid")
	if err := user.FindUserByID(userid); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err := utils.DeleteFile(user.Iconpath)
	if err != nil && user.Iconpath != "" {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	outputDir := utils.ProjectRoot + "/storage/users/" + userid
	outputFile := utils.ProjectRoot + "/storage/users/" + userid + "/"+ uuid.New().String()
	os.Mkdir(outputDir, 0777)
	out, err := os.Create(outputFile)
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	user.Iconpath = outputFile
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
	userid := c.Param("userid")
	if err := user.FindUserByID(userid); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	
	if err := utils.DeleteFile(user.Iconpath); err != nil {
		c.String(http.StatusNotFound, "Not Found")
        return
	}
	user.Iconpath = ""
	if err := user.UpdateUserByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func GetUserIcon(c *gin.Context) {
	var user models.User
	userid := c.Param("userid")
	if err := user.FindUserByID(userid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if user.Iconpath == "" {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	c.File(user.Iconpath)
}

// func UploadUserIcon(c *gin.Context) {
// 	var user models.User
// 	userid := c.Param("userid")
// 	if err := user.FindUserByID(userid); err == exceptions.NotFound {
// 		c.String(http.StatusBadRequest, "Bad request")
// 		return
// 	}
// 	// err := utils.DeleteFile(user.Iconpath)
// 	if err != nil && user.Iconpath != "" {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 		return
// 	}
// 	var image utils.Icon
// 	if err := c.ShouldBindJSON(&image); err != nil {
// 		c.String(http.StatusBadRequest, "Bad request")
// 		return
// 	}
// 	outputFile := utils.ProjectRoot + "/storage/users/" + userid
// 	savePath, err := utils.SaveDecodedIcon(image.EncodedData, outputFile)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
//         return
// 	}
// 	user.Iconpath = savePath
// 	if err := user.UpdateUserByID(); err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
//         return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"user": user,
// 	})
// }

// func DeleteUserIcon(c *gin.Context) {
// 	var user models.User
// 	userid := c.Param("userid")
// 	if err := user.FindUserByID(userid); err == exceptions.NotFound {
// 		c.String(http.StatusNotFound, "Not Found")
// 		return
// 	}
	
// 	if err := utils.DeleteFile(user.Iconpath); err != nil {
// 		c.String(http.StatusNotFound, "Not Found")
//         return
// 	}
// 	user.Iconpath = ""
// 	if err := user.UpdateUserByID(); err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
//         return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"user": user,
// 	})
// }