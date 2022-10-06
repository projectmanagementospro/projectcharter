package service

import (
	"fmt"
	"projectcharter/models/domain"
	"projectcharter/models/web"
	"projectcharter/repository"

	"github.com/mashingan/smapping"
)

type ProjectCharterService interface {
	All() []domain.ProjectCharter
	Create(request web.ProjectCharterRequest) (domain.ProjectCharter, error)
	FindById(id uint) (domain.ProjectCharter, error)
	Update(request web.ProjectCharterUpdateRequest) (domain.ProjectCharter, error)
	Delete(id uint) error
}

type projectcharterService struct {
	projectcharterRepository repository.ProjectCharterRepository
}

func NewProjectCharterService(projectcharterRepository repository.ProjectCharterRepository) ProjectCharterService {
	return &projectcharterService{projectcharterRepository: projectcharterRepository}
}

func (projectcharterservice *projectcharterService) All() []domain.ProjectCharter {
	return projectcharterservice.projectcharterRepository.All()
}

func (projectcharterservice *projectcharterService) Create(request web.ProjectCharterRequest) (domain.ProjectCharter, error) {
	projectcharter := domain.ProjectCharter{}

	//time.Date(request.StartDate.Year(), request.StartDate.Month(), request.StartDate.Day(), 0, 0, 0, 0, time.Local)
	//projectcharter.StartDate = utils.ConvertDate(request.StartDate)
	//projectcharter.EndDate = utils.ConvertDate(request.EndDate)
	//request.StartDate = nil
	//request.EndDate = nil

	err := smapping.FillStruct(&projectcharter, smapping.MapFields(&request))
	if err != nil {

		return projectcharter, err
	}
	fmt.Println(projectcharter)
	// _, err = projectcharterservice.projectcharterRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return projectcharter, err
	// }
	return projectcharterservice.projectcharterRepository.Create(projectcharter), nil
}

func (projectcharterservice *projectcharterService) Update(request web.ProjectCharterUpdateRequest) (domain.ProjectCharter, error) {
	projectcharter := domain.ProjectCharter{}
	res, err := projectcharterservice.projectcharterRepository.FindById(request.ID)
	if err != nil {
		return projectcharter, err
	}
	err = smapping.FillStruct(&projectcharter, smapping.MapFields(&request))
	if err != nil {
		return projectcharter, err
	}
	//projectcharter.ID = res.ID
	projectcharter.User_id = res.User_id
	return projectcharterservice.projectcharterRepository.Update(projectcharter), nil
}

func (projectcharterservice *projectcharterService) FindById(id uint) (domain.ProjectCharter, error) {
	projectcharter, err := projectcharterservice.projectcharterRepository.FindById(id)
	if err != nil {
		return projectcharter, err
	}
	return projectcharter, nil
}

func (projectcharterservice *projectcharterService) Delete(id uint) error {
	projectcharter, err := projectcharterservice.projectcharterRepository.FindById(id)
	if err != nil {
		return err
	}
	projectcharterservice.projectcharterRepository.Delete(projectcharter)
	return nil
}
