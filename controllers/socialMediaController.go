package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go-final-mygram/database"
	"go-final-mygram/helpers"
	"go-final-mygram/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllSocialMedia(c *gin.Context) {
	var (
		photos []models.Photo
	)

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserId = userID

	if err := db.Model(&models.Photo{}).Order("created_at asc").Find(&photos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Getting Order Data: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": photos,
	})
}

func GetOneSocialMedia(c *gin.Context) {
	var (
		photoId string
		photo   models.Photo
	)
	db := database.GetDB()
	photoId = c.Param("photoId")

	if err := db.First(&models.Photo{}, "id = ?", photoId).Find(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Photo with id %v Not Found", photoId),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data found",
		"data":    photo,
	})
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserId = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userId
	Photo.ID = uint(photoId)
	err := db.Model(&Photo).Where("id=?", photoId).Updates(models.Photo{UserId: userId, Url: Photo.Url}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

func DeleteSocialMedia(c *gin.Context) {
	var (
		photoId string
	)
	db := database.GetDB()
	photoId = c.Param("photoId")

	if err := db.First(&models.Photo{}, "id = ?", photoId).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Photo with id %v Not Found", photoId),
		})
		return
	}

	if err := db.Where("id = ?", photoId).Delete(&models.Photo{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Photo: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with id %v Has Been Successfully Deleted", photoId),
		"success": true,
	})
}
