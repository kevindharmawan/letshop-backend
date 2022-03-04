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
			profile.GET("/get", api.GetProfile)
			profile.GET("/get/:id", api.GetProfileById)
			profile.POST("/create", api.CreateProfile)
			profile.PATCH("/update", api.UpdateProfile)
			profile.DELETE("/delete", api.DeleteProfile)
		}
	}

	return router
}
