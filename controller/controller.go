package controller

import (
	"projectcharter/config"
	"projectcharter/models/domain"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CreatedBy   string `json:"created_by"`
		UpdatedBy   string `json:"updated_by"`
		DeletedBy   string `json:"deleted_by"`
	}

	c.BindJSON(&body)

	// Create a post
	projectcharter := domain.ProjectCharter{Name: body.Name, Description: body.Description, CreatedBy: body.CreatedBy, UpdatedBy: body.UpdatedBy, DeletedBy: body.DeletedBy}

	result := config.NewDB().Create(&projectcharter)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"projectcharter": projectcharter,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post domain.ProjectCharter
	config.NewDB().First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}
