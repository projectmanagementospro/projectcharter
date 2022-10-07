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

func (pCharterService *projectcharterService) All() []domain.ProjectCharter {
	return pCharterService.projectcharterRepository.All()
}

func (pCharterService *projectcharterService) Create(request web.ProjectCharterRequest) (domain.ProjectCharter, error) {
	pCharter := domain.ProjectCharter{}

	//time.Date(request.StartDate.Year(), request.StartDate.Month(), request.StartDate.Day(), 0, 0, 0, 0, time.Local)
	//pCharter.StartDate = utils.ConvertDate(request.StartDate)
	//pCharter.EndDate = utils.ConvertDate(request.EndDate)
	//request.StartDate = nil
	//request.EndDate = nil

	err := smapping.FillStruct(&pCharter, smapping.MapFields(&request))
	if err != nil {

		return pCharter, err
	}
	fmt.Println(pCharter)
	// _, err = pCharterService.projectcharterRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return pCharter, err
	// }
	return pCharterService.projectcharterRepository.Create(pCharter), nil
}

func (pCharterService *projectcharterService) Update(request web.ProjectCharterUpdateRequest) (domain.ProjectCharter, error) {
	pCharter := domain.ProjectCharter{}
	res, err := pCharterService.projectcharterRepository.FindById(request.ID)
	if err != nil {
		return pCharter, err
	}
	err = smapping.FillStruct(&pCharter, smapping.MapFields(&request))
	if err != nil {
		return pCharter, err
	}
	//pCharter.ID = res.ID
	pCharter.User_id = res.User_id
	return pCharterService.projectcharterRepository.Update(pCharter), nil
}

func (pCharterService *projectcharterService) FindById(id uint) (domain.ProjectCharter, error) {
	pCharter, err := pCharterService.projectcharterRepository.FindById(id)
	if err != nil {
		return pCharter, err
	}
	return pCharter, nil
}

func (pCharterService *projectcharterService) Delete(id uint) error {
	pCharter, err := pCharterService.projectcharterRepository.FindById(id)
	if err != nil {
		return err
	}
	pCharterService.projectcharterRepository.Delete(pCharter)
	return nil
}
