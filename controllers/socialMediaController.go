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
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}

	err := db.Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedias)
}

func GetOneSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.First(&SocialMedia, "id = ?", socialmediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	SocialMedia.UserId = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.First(&SocialMedia, "id = ?", socialmediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Delete(&SocialMedia)

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("socialmedia with id %v has been successfully deleted", socialmediaId),
	})
}
