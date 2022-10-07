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

func (pCharterController *projectcharterController) All(context *gin.Context) {
	pcharters := pCharterController.projectcharterService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pcharters,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (pCharterController *projectcharterController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	pcharter, err := pCharterController.projectcharterService.FindById(uint(id))
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

func (pCharterController *projectcharterController) Insert(context *gin.Context) {
	var request web.ProjectCharterRequest
	err := context.BindJSON(&request)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	request.User_id = 1
	pcharter, err := pCharterController.projectcharterService.Create(request)
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

func (pCharterController *projectcharterController) Update(context *gin.Context) {
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
	pcharter, err := pCharterController.projectcharterService.Update(request)
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

func (pCharterController *projectcharterController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = pCharterController.projectcharterService.Delete(uint(id))
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

// func ProjCharterCreate(pCharterController *gin.Context) {
// 	// Get data off req body
// 	var body struct {
// 		Name        string `json:"name"`
// 		Description string `json:"description"`
// 		CreatedBy   string `json:"created_by"`
// 		UpdatedBy   string `json:"updated_by"`
// 		DeletedBy   string `json:"deleted_by"`
// 	}

// 	pCharterController.BindJSON(&body)

// 	// Create a post
// 	projectcharter := domain.ProjectCharter{Name: body.Name, Description: body.Description, CreatedBy: body.CreatedBy, UpdatedBy: body.UpdatedBy, DeletedBy: body.DeletedBy}

// 	result := config.NewDB().Create(&projectcharter)

// 	if result.Error != nil {
// 		pCharterController.Status(400)
// 		return
// 	}

// 	// Return it
// 	pCharterController.JSON(200, gin.H{
// 		"projectcharter": projectcharter,
// 	})
// }

// func ProjCharterIndex(pCharterController *gin.Context) {
// 	// Get the pcharter
// 	var pcharter []domain.ProjectCharter
// 	config.NewDB().Find(&pcharter)

// 	// Respond with them
// 	pCharterController.JSON(200, gin.H{
// 		"pcharter": pcharter,
// 	})
// }

// func ProjCharterShow(pCharterController *gin.Context) {
// 	// Get id off url
// 	id := pCharterController.Param("id")

// 	// Get the posts
// 	var projectcharter domain.ProjectCharter
// 	config.NewDB().First(&projectcharter, id)

// 	// Respond with them
// 	pCharterController.JSON(200, gin.H{
// 		"pcharter": projectcharter,
// 	})
// }

// func ProjCharterUpdate(pCharterController *gin.Context) {
// 	// Get the id off the url
// 	id := pCharterController.Param("id")

// 	// Get the data off req body
// 	var body struct {
// 		Name        string `json:"name"`
// 		Description string `json:"description"`
// 		CreatedBy   string `json:"created_by"`
// 		UpdatedBy   string `json:"updated_by"`
// 		DeletedBy   string `json:"deleted_by"`
// 	}

// 	pCharterController.Bind(&body)

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
// 	pCharterController.JSON(200, gin.H{
// 		"pcharter": projectcharter,
// 	})
// }

// func ProjCharterDelete(pCharterController *gin.Context) {
// 	// Get the id off the url
// 	id := pCharterController.Param("id")

// 	// Delete the pcharter
// 	config.NewDB().Delete(&domain.ProjectCharter{}, id)

// 	// Respond
// 	pCharterController.Status(200)
// }
