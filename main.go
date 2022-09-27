package main

import (
	"projectcharter/config"
	"projectcharter/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	config.NewDB()
}

func main() {

	r := gin.Default()

	r.POST("/pcharter", controller.ProjCharterCreate)
	r.POST("/pcharter/:id", controller.ProjCharterUpdate)

	r.GET("/pcharter/:id", controller.ProjCharterShow)
	r.GET("/pcharter", controller.ProjCharterIndex)

	r.DELETE("/pcharter/:id", controller.ProjCharterDelete)

	r.Run("0.0.0.0:3000")
	//r.Run()
}
