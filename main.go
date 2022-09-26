package main

import (
	"projectcharter/controller"
	"projectcharter/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controller.PostsCreate)

	r.GET("/posts/:id", controller.PostsShow)

	r.Run("0.0.0.0:3000")
	//r.Run()
}
