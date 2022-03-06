package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"letshop-backend/constants"
	"letshop-backend/models"
	"letshop-backend/services"
)

const notFoundMessage = "profile not found"
const updateErrorMessage = "error when updating profile"

func CreateProfile(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var input models.ProfileCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	// TODO: Implement image
	profile := models.Profile{
		UserID:    uuid,
		Name:      input.Name,
		BirthDate: input.BirthDate,
		ImageUrl:  "",
		Gender:    input.Gender,
	}
	db.Create(&profile)

	c.JSON(http.StatusOK, models.Response{Data: profile})
}

func UpdateProfile(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var input models.ProfileUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	var profile models.Profile
	if result := db.Model(&models.Profile{}).Where("user_id = ?", uuid).First(&profile); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: notFoundMessage})
		return
	}

	if result := db.Model(&profile).Updates(models.Profile{
		Name:      input.Name,
		BirthDate: input.BirthDate,
		Gender:    input.Gender,
	}); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: updateErrorMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: profile})
}

func GetMyProfile(c *gin.Context) {
	db := services.Database

	uuid := c.Query(constants.UserIDKey)
	if uuid == "" {
		uuid = c.GetString(constants.UserIDKey)
	}

	var profile models.Profile
	if result := db.Model(&models.Profile{}).Where("user_id = ?", uuid).First(&profile); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: notFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: profile})
}

// TODO: Change to soft delete
func DeleteProfile(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	if result := db.Delete(&models.Profile{}, "user_id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: notFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: "Profile deleted"})
}
