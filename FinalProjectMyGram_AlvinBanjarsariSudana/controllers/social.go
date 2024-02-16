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

func CreateSocial(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Social := models.Social{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.UserID = userID

	err := db.Debug().Create(&Social).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Social)
}

func UpdateSocial(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Social := models.Social{}

	socialId, _ := strconv.Atoi(c.Param("socialId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.UserID = userID
	Social.ID = uint(socialId)

	err := db.Model(&Social).Where("id = ?", socialId).Updates(models.Social{Name: Social.Name, Social_url: Social.Social_url}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Social)
}

func GetAllSocial(c *gin.Context) {
	db := database.GetDB()

	social := []models.Social{}

	err := db.Debug().Preload("User").Order("id asc").Find(&social).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaString, _ := json.Marshal(social)
	socialMediaResponse := []models.Social{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusOK, socialMediaResponse)
}

func GetByIdSocial(c *gin.Context) {
	db := database.GetDB()
	socialId := c.Param("socialId")
	var allSocial []models.Social

	err := db.First(&allSocial, "Id = ?", socialId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": allSocial,
	})
}

func DeleteSocial(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialId, _ := strconv.Atoi(c.Param("socialId"))
	userID := uint(userData["id"].(float64))

	social := models.Social{}
	social.ID = uint(socialId)
	social.UserID = userID

	err := db.Delete(&social).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
