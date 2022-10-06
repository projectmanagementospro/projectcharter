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

func (projectchartercontroller *projectcharterController) All(context *gin.Context) {
	pcharters := projectchartercontroller.projectcharterService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pcharters,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (projectchartercontroller *projectcharterController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	pcharter, err := projectchartercontroller.projectcharterService.FindById(uint(id))
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

func (projectchartercontroller *projectcharterController) Insert(context *gin.Context) {
	var request web.ProjectCharterRequest
	err := context.BindJSON(&request)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	request.User_id = 1
	pcharter, err := projectchartercontroller.projectcharterService.Create(request)
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

func (projectchartercontroller *projectcharterController) Update(context *gin.Context) {
	var request web.ProjectCharterUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	request.ID = uint(id)
	err = context.BindJSON(&request)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	pcharter, err := projectchartercontroller.projectcharterService.Update(request)
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

func (projectchartercontroller *projectcharterController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = projectchartercontroller.projectcharterService.Delete(uint(id))
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

// func ProjCharterCreate(projectchartercontroller *gin.Context) {
// 	// Get data off req body
// 	var body struct {
// 		Name        string `json:"name"`
// 		Description string `json:"description"`
// 		CreatedBy   string `json:"created_by"`
// 		UpdatedBy   string `json:"updated_by"`
// 		DeletedBy   string `json:"deleted_by"`
// 	}

// 	projectchartercontroller.BindJSON(&body)

// 	// Create a post
// 	projectcharter := domain.ProjectCharter{Name: body.Name, Description: body.Description, CreatedBy: body.CreatedBy, UpdatedBy: body.UpdatedBy, DeletedBy: body.DeletedBy}

// 	result := config.NewDB().Create(&projectcharter)

// 	if result.Error != nil {
// 		projectchartercontroller.Status(400)
// 		return
// 	}

// 	// Return it
// 	projectchartercontroller.JSON(200, gin.H{
// 		"projectcharter": projectcharter,
// 	})
// }

// func ProjCharterIndex(projectchartercontroller *gin.Context) {
// 	// Get the pcharter
// 	var pcharter []domain.ProjectCharter
// 	config.NewDB().Find(&pcharter)

// 	// Respond with them
// 	projectchartercontroller.JSON(200, gin.H{
// 		"pcharter": pcharter,
// 	})
// }

// func ProjCharterShow(projectchartercontroller *gin.Context) {
// 	// Get id off url
// 	id := projectchartercontroller.Param("id")

// 	// Get the posts
// 	var projectcharter domain.ProjectCharter
// 	config.NewDB().First(&projectcharter, id)

// 	// Respond with them
// 	projectchartercontroller.JSON(200, gin.H{
// 		"pcharter": projectcharter,
// 	})
// }

// func ProjCharterUpdate(projectchartercontroller *gin.Context) {
// 	// Get the id off the url
// 	id := projectchartercontroller.Param("id")

// 	// Get the data off req body
// 	var body struct {
// 		Name        string `json:"name"`
// 		Description string `json:"description"`
// 		CreatedBy   string `json:"created_by"`
// 		UpdatedBy   string `json:"updated_by"`
// 		DeletedBy   string `json:"deleted_by"`
// 	}

// 	projectchartercontroller.Bind(&body)

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
// 	projectchartercontroller.JSON(200, gin.H{
// 		"pcharter": projectcharter,
// 	})
// }

// func ProjCharterDelete(projectchartercontroller *gin.Context) {
// 	// Get the id off the url
// 	id := projectchartercontroller.Param("id")

// 	// Delete the pcharter
// 	config.NewDB().Delete(&domain.ProjectCharter{}, id)

// 	// Respond
// 	projectchartercontroller.Status(200)
// }
