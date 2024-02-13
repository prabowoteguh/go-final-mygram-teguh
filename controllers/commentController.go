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

func GetAllComment(c *gin.Context) {
	var (
		comments []models.Comment
	)

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserId = userID

	if err := db.Model(&models.Comment{}).Order("created_at asc").Find(&comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Getting Order Data: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": comments,
	})
}

func GetOneComment(c *gin.Context) {
	var (
		commentId string
		comment   models.Comment
	)
	db := database.GetDB()
	commentId = c.Param("commentId")

	if err := db.First(&models.Comment{}, "id = ?", commentId).Find(&comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Comment with id %v Not Found", commentId),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data found",
		"Comment": comment,
	})
}

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserId = userID
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	CommentId, _ := strconv.Atoi(c.Param("CommentId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userId
	Comment.ID = uint(CommentId)
	err := db.Model(&Comment).Where("id=?", CommentId).Updates(models.Comment{UserId: userId, Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	var (
		commentId string
	)
	db := database.GetDB()
	commentId = c.Param("commentId")

	if err := db.First(&models.Comment{}, "id = ?", commentId).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Comment with id %v Not Found", commentId),
		})
		return
	}

	if err := db.Where("id = ?", commentId).Delete(&models.Comment{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Comment: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Comment with id %v Has Been Successfully Deleted", commentId),
		"success": true,
	})
}
