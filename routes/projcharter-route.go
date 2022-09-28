package routes

import (
	"projectcharter/injector"
	"projectcharter/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProjectCharterRoutes(db *gorm.DB, route *gin.Engine) {
	pcharterController := injector.InitProjectCharter(db)
	pcharterRoute := route.Group("/api/v1/pcharter")
	pcharterRoute.Use(middleware.ErrorHandler())
	pcharterRoute.Use(cors.Default())
	pcharterRoute.GET("/", pcharterController.All)
	pcharterRoute.GET("/:id", pcharterController.FindById)
	pcharterRoute.POST("/", pcharterController.Insert)
	pcharterRoute.PUT("/:id", pcharterController.Update)
	pcharterRoute.DELETE("/:id", pcharterController.Delete)
}
