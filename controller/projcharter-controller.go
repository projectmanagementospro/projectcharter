package controller

import (
	"net/http"
	"projectcharter/helper"
	"projectcharter/models/web"
	"projectcharter/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectCharterController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type projectcharterController struct {
	projectcharterService service.ProjectCharterService
}

func NewProjectCharterController(projectcharterService service.ProjectCharterService) ProjectCharterController {
	return &projectcharterController{
		projectcharterService: projectcharterService,
	}
}

func (c *projectcharterController) All(context *gin.Context) {
	pcharters := c.projectcharterService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pcharters,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *projectcharterController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	pcharter, err := c.projectcharterService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pcharter,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *projectcharterController) Insert(context *gin.Context) {
	var u web.ProjectCharterRequest
	err := context.BindJSON(&u)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	u.User_id = 1
	pcharter, err := c.projectcharterService.Create(u)
	println("ada")
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	println("tidak ada")
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pcharter,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *projectcharterController) Update(context *gin.Context) {
	var u web.ProjectCharterUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	u.ID = uint(id)
	err = context.BindJSON(&u)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	pcharter, err := c.projectcharterService.Update(u)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pcharter,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *projectcharterController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = c.projectcharterService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   "Project charter has been removed",
	}
	context.JSON(http.StatusOK, webResponse)
}

// func ProjCharterCreate(c *gin.Context) {
// 	// Get data off req body
// 	var body struct {
// 		Name        string `json:"name"`
// 		Description string `json:"description"`
// 		CreatedBy   string `json:"created_by"`
// 		UpdatedBy   string `json:"updated_by"`
// 		DeletedBy   string `json:"deleted_by"`
// 	}

// 	c.BindJSON(&body)

// 	// Create a post
// 	projectcharter := domain.ProjectCharter{Name: body.Name, Description: body.Description, CreatedBy: body.CreatedBy, UpdatedBy: body.UpdatedBy, DeletedBy: body.DeletedBy}

// 	result := config.NewDB().Create(&projectcharter)

// 	if result.Error != nil {
// 		c.Status(400)
// 		return
// 	}

// 	// Return it
// 	c.JSON(200, gin.H{
// 		"projectcharter": projectcharter,
// 	})
// }

// func ProjCharterIndex(c *gin.Context) {
// 	// Get the pcharter
// 	var pcharter []domain.ProjectCharter
// 	config.NewDB().Find(&pcharter)

// 	// Respond with them
// 	c.JSON(200, gin.H{
// 		"pcharter": pcharter,
// 	})
// }

// func ProjCharterShow(c *gin.Context) {
// 	// Get id off url
// 	id := c.Param("id")

// 	// Get the posts
// 	var projectcharter domain.ProjectCharter
// 	config.NewDB().First(&projectcharter, id)

// 	// Respond with them
// 	c.JSON(200, gin.H{
// 		"pcharter": projectcharter,
// 	})
// }

// func ProjCharterUpdate(c *gin.Context) {
// 	// Get the id off the url
// 	id := c.Param("id")

// 	// Get the data off req body
// 	var body struct {
// 		Name        string `json:"name"`
// 		Description string `json:"description"`
// 		CreatedBy   string `json:"created_by"`
// 		UpdatedBy   string `json:"updated_by"`
// 		DeletedBy   string `json:"deleted_by"`
// 	}

// 	c.Bind(&body)

// 	// Find the post were updating
// 	var projectcharter domain.ProjectCharter
// 	config.NewDB().First(&projectcharter, id)

// 	// Update it
// 	config.NewDB().Model(&projectcharter).Updates(domain.ProjectCharter{
// 		Name:        body.Name,
// 		Description: body.Description,
// 		CreatedBy:   body.CreatedBy,
// 		UpdatedBy:   body.UpdatedBy,
// 		DeletedBy:   body.DeletedBy,
// 	})

// 	// Respond with it
// 	c.JSON(200, gin.H{
// 		"pcharter": projectcharter,
// 	})
// }

// func ProjCharterDelete(c *gin.Context) {
// 	// Get the id off the url
// 	id := c.Param("id")

// 	// Delete the pcharter
// 	config.NewDB().Delete(&domain.ProjectCharter{}, id)

// 	// Respond
// 	c.Status(200)
// }
