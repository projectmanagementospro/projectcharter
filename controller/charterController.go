package controller

import (
	"projectcharter/config"
	"projectcharter/models/domain"

	"github.com/gin-gonic/gin"
)

func ProjCharterCreate(c *gin.Context) {
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

func ProjCharterIndex(c *gin.Context) {
	// Get the pcharter
	var pcharter []domain.ProjectCharter
	config.NewDB().Find(&pcharter)

	// Respond with them
	c.JSON(200, gin.H{
		"pcharter": pcharter,
	})
}

func ProjCharterShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var projectcharter domain.ProjectCharter
	config.NewDB().First(&projectcharter, id)

	// Respond with them
	c.JSON(200, gin.H{
		"pcharter": projectcharter,
	})
}

func ProjCharterUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CreatedBy   string `json:"created_by"`
		UpdatedBy   string `json:"updated_by"`
		DeletedBy   string `json:"deleted_by"`
	}

	c.Bind(&body)

	// Find the post were updating
	var projectcharter domain.ProjectCharter
	config.NewDB().First(&projectcharter, id)

	// Update it
	config.NewDB().Model(&projectcharter).Updates(domain.ProjectCharter{
		Name:        body.Name,
		Description: body.Description,
		CreatedBy:   body.CreatedBy,
		UpdatedBy:   body.UpdatedBy,
		DeletedBy:   body.DeletedBy,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"pcharter": projectcharter,
	})
}

func ProjCharterDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the pcharter
	config.NewDB().Delete(&domain.ProjectCharter{}, id)

	// Respond
	c.Status(200)
}
