package router

import (
	"github.com/gin-gonic/gin"

	"letshop-backend/controller/api"
	"letshop-backend/controller/middleware"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	apiRoute := router.Group("/api")
	apiRoute.Use(
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
	)
	{
		profile := apiRoute.Group("/profile")
		{
			profile.GET("/get", api.GetMyProfile)
			profile.POST("/create", api.CreateProfile)
			profile.PATCH("/update", api.UpdateProfile)
			profile.DELETE("/delete", api.DeleteProfile)
		}
		shop := apiRoute.Group("/shop")
		{
			shop.GET("/get", api.GetMyShop)
			shop.GET("/get/:id", api.GetShopById)
			shop.POST("/create", api.CreateShop)
			shop.PATCH("/update", api.UpdateShop)
			shop.DELETE("/delete", api.DeleteShop)
		}
		product := apiRoute.Group("/product")
		{
			product.GET("/get", api.GetMyProducts)
			product.GET("/get/recommended", api.GetRecommendedProducts)
			product.GET("/get/:id", api.GetProductById)
			// product.GET("/get/:category_id", api.GetProductsByCategory)
			// product.GET("/get/:shop_id", api.GetProductsByShop)
			product.POST("/create", api.CreateProduct)
			product.PATCH("/update", api.UpdateProduct)
			product.DELETE("/delete", api.DeleteProduct)
		}
		category := apiRoute.Group("/category")
		{
			category.GET("/get/recommended", api.GetRecommendedCategories)
			category.GET("/get/:id", api.GetCategoryById)
			category.POST("/create", api.CreateCategory)
			category.PATCH("/update", api.UpdateCategory)
			category.DELETE("/delete", api.DeleteCategory)
		}
		wishlist := apiRoute.Group("/wishlist")
		{
			wishlist.GET("/get", api.GetMyWishlists)
			wishlist.POST("/create", api.CreateWishlist)
			wishlist.DELETE("/delete", api.DeleteWishlist)
		}
	}

	return router
}
