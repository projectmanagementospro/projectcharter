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

	r.POST("/posts", controller.PostsCreate)

	r.GET("/posts/:id", controller.PostsShow)

	r.Run("0.0.0.0:3000")
	//r.Run()
}
