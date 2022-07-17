package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"letshop-backend/constants"
	"letshop-backend/models"
	"letshop-backend/services"
)

const shopNotFoundMessage = "shop not found"
const shopUpdateErrorMessage = "error when updating shop"

func CreateShop(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	// TODO: Check if user already have shop

	var input models.ShopCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	// TODO: Implement image
	shop := models.Shop{
		UserID:   uuid,
		Name:     input.Name,
		Desc:     input.Desc,
		ImageUrl: "",
	}
	db.Create(&shop)

	c.JSON(http.StatusOK, models.Response{Data: shop})
}

func UpdateShop(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var input models.ShopUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	var shop models.Shop
	if result := db.Model(&models.Shop{}).Where("user_id = ?", uuid).First(&shop); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: shopNotFoundMessage})
		return
	}

	if result := db.Model(&shop).Updates(models.Shop{
		Name:     input.Name,
		Desc:     input.Desc,
		ImageUrl: "",
	}); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: shopUpdateErrorMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: shop})
}

func GetMyShop(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var shop models.Shop
	if result := db.Preload("Products").First(&shop, "user_id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: shopNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: shop})
}

func GetShopById(c *gin.Context) {
	db := services.Database

	var shop models.Shop
	if result := db.Preload("Products").First(&shop, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: shopNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: shop})
}

func DeleteShop(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	if result := db.Delete(&models.Shop{}, "user_id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: shopNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: "Shop deleted"})
}
