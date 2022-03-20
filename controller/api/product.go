package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"letshop-backend/constants"
	"letshop-backend/models"
	"letshop-backend/services"
)

const productNotFoundMessage = "product not found"
const productNoShopMessage = "you need to create a shop first"
const productLoadErrorMessage = "error when loading product"
const productUpdateErrorMessage = "error when updating product"

func CreateProduct(c *gin.Context) {
	db := services.Database

	uuid := c.GetString(constants.UserIDKey)

	var shop models.Shop
	if result := db.First(&shop, "user_id = ?", uuid); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: productNoShopMessage})
		return
	}

	var input models.ProductCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	// TODO: Implement image
	product := models.Product{
		Name:      input.Name,
		Desc:      input.Desc,
		Price:     input.Price,
		ShopID:    shop.ID,
		Stock:     input.Stock,
		Discount:  input.Discount,
		Thumbnail: "",
	}
	db.Create(&product)

	c.JSON(http.StatusOK, models.Response{Data: product})
}

// TODO: Implement
func UpdateProduct(c *gin.Context) { return }

// TODO: Implement
func GetMyProducts(c *gin.Context) { return }

func GetRecommendedProducts(c *gin.Context) {
	db := services.Database

	var products []models.Product
	if result := db.Find(&products); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: productLoadErrorMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: products})
}

func GetProductById(c *gin.Context) {
	db := services.Database

	var product models.Product
	if result := db.Preload("Images").First(&product, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: productNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: product})
}

// TODO: Implement
func GetProductsByCategory(c *gin.Context) { return }

// TODO: Implement
func GetProductsByShop(c *gin.Context) { return }

// TODO: Implement
func DeleteProduct(c *gin.Context) {
	// db := services.Database

	// uuid := c.GetString(constants.UserIDKey)

	// var product models.Product
	// db.Preload("")

	// if result := db.Delete(&models.Shop{}, "user_id = ?", uuid); result.Error != nil {
	// 	c.JSON(http.StatusNotFound, models.Response{Error: shopNotFoundMessage})
	// 	return
	// }

	// c.JSON(http.StatusOK, models.Response{Data: "Shop deleted"})
}
