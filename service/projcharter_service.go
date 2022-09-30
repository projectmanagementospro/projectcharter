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
	Create(b web.ProjectCharterRequest) (domain.ProjectCharter, error)
	FindById(id uint) (domain.ProjectCharter, error)
	Update(b web.ProjectCharterUpdateRequest) (domain.ProjectCharter, error)
	Delete(id uint) error
}

type projectcharterService struct {
	projectcharterRepository repository.ProjectCharterRepository
}

func NewProjectCharterService(projectcharterRepository repository.ProjectCharterRepository) ProjectCharterService {
	return &projectcharterService{projectcharterRepository: projectcharterRepository}
}

func (s *projectcharterService) All() []domain.ProjectCharter {
	return s.projectcharterRepository.All()
}

func (s *projectcharterService) Create(request web.ProjectCharterRequest) (domain.ProjectCharter, error) {
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
	// _, err = s.projectcharterRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return projectcharter, err
	// }
	return s.projectcharterRepository.Create(projectcharter), nil
}

func (s *projectcharterService) Update(b web.ProjectCharterUpdateRequest) (domain.ProjectCharter, error) {
	projectcharter := domain.ProjectCharter{}
	res, err := s.projectcharterRepository.FindById(b.ID)
	if err != nil {
		return projectcharter, err
	}
	err = smapping.FillStruct(&projectcharter, smapping.MapFields(&b))
	if err != nil {
		return projectcharter, err
	}
	//projectcharter.ID = res.ID
	projectcharter.User_id = res.User_id
	return s.projectcharterRepository.Update(projectcharter), nil
}

func (s *projectcharterService) FindById(id uint) (domain.ProjectCharter, error) {
	projectcharter, err := s.projectcharterRepository.FindById(id)
	if err != nil {
		return projectcharter, err
	}
	return projectcharter, nil
}

func (s *projectcharterService) Delete(id uint) error {
	projectcharter, err := s.projectcharterRepository.FindById(id)
	if err != nil {
		return err
	}
	s.projectcharterRepository.Delete(projectcharter)
	return nil
}
