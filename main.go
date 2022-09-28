package main

import (
	"projectcharter/config"
	"projectcharter/helper"
	"projectcharter/injector"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.NewDB()
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	// pcharterRepository := repository.NewProjectCharterRepository(db)
	// pcharterService := service.NewProjectCharterService(pcharterRepository)
	// pcharterController := controller.NewProjectCharterController(pcharterService)

	pcharterController := injector.InitProjectCharter(db)

	r := gin.Default()
	pcharter := r.Group("/pcharter")

	pcharter.GET("/", pcharterController.All)
	pcharter.GET("/:id", pcharterController.FindById)
	pcharter.POST("/", pcharterController.Insert)
	pcharter.PUT("/:id", pcharterController.Update)
	pcharter.DELETE("/:id", pcharterController.Delete)

	r.Run("0.0.0.0:8080")

	// router := NewRouter()
	// log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

// func NewRouter() *gin.Engine {
// 	err := godotenv.Load()
// 	helper.PanicIfError(err)
// 	/**
// 	@description Setup Database Connection
// 	*/

// 	/**
// 	@description Init Router
// 	*/
// 	router := gin.Default()
// 	/**
// 	@description Setup Mode Application
// 	*/
// 	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
// 		gin.SetMode(gin.DebugMode)
// 	} else if os.Getenv("GO_ENV") == "test" {
// 		gin.SetMode(gin.TestMode)
// 	} else {
// 		gin.SetMode(gin.ReleaseMode)
// 	}
// 	/**
// 	@description Setup Middleware
// 	*/

// 	/**
// 	@description Init All Route
// 	*/
// 	routes.NewAuthenticationRoutes(db, router)
// 	routes.NewUserRoutes(db, router)
// 	router.Use(middleware.ErrorHandler())
// 	router.Use(cors.Default())

// 	return router
// }
