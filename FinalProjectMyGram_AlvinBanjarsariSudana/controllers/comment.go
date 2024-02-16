package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Photo.ID = uint(photoId)
	Comment.PhotoID = Photo.ID

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

	commentRequest := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))
	//Pengecekan content-type yang direquest melalui API
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserID = userID

	updateString, _ := json.Marshal(commentRequest)
	updateData := models.Comment{}
	json.Unmarshal(updateString, &updateData)
	err := db.Model(&comment).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&comment, comment.ID).Error

	commentString, _ := json.Marshal(comment)
	commentResponse := models.Comment{}
	json.Unmarshal(commentString, &commentResponse)

	c.JSON(http.StatusOK, commentResponse)
}

func GetAllComment(c *gin.Context) {
	db := database.GetDB()

	comments := []models.Comment{}

	err := db.Debug().Preload("User").Preload("Photo").Order("id asc").Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	commentsString, _ := json.Marshal(comments)
	commentsResponse := []models.Comment{}
	json.Unmarshal(commentsString, &commentsResponse)

	c.JSON(http.StatusOK, commentsResponse)
}

func GetByIdComment(c *gin.Context) {
	db := database.GetDB()
	commentId := c.Param("commentId")
	var commentById []models.Comment

	err := db.First(&commentById, "Id = ?", commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record no found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Comment": commentById,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserID = userID

	err := db.Delete(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
