package repository

import (
	"errors"
	"projectcharter/models/domain"

	"gorm.io/gorm"
)

type ProjectCharterRepository interface {
	All() []domain.ProjectCharter
	Create(pCharter domain.ProjectCharter) domain.ProjectCharter
	Update(pCharter domain.ProjectCharter) domain.ProjectCharter
	Delete(pCharter domain.ProjectCharter)
	FindById(id uint) (domain.ProjectCharter, error)
}

type ProjectCharterConnection struct {
	//Connection to database
	dbConnect *gorm.DB
}

func NewProjectCharterRepository(db *gorm.DB) ProjectCharterRepository {
	return &ProjectCharterConnection{dbConnect: db} //
}

func (conn *ProjectCharterConnection) All() []domain.ProjectCharter {
	var projectcharters []domain.ProjectCharter
	conn.dbConnect.Find(&projectcharters)
	return projectcharters
}

func (conn *ProjectCharterConnection) Create(pCharter domain.ProjectCharter) domain.ProjectCharter {
	conn.dbConnect.Save(&pCharter)
	return pCharter
}

func (conn *ProjectCharterConnection) Update(pCharter domain.ProjectCharter) domain.ProjectCharter {
	conn.dbConnect.Omit("created_at").Save(&pCharter)
	return pCharter
}

func (conn *ProjectCharterConnection) Delete(pCharter domain.ProjectCharter) {
	conn.dbConnect.Delete(&pCharter)
}

func (conn *ProjectCharterConnection) FindById(id uint) (domain.ProjectCharter, error) {
	var pCharter domain.ProjectCharter
	conn.dbConnect.Find(&pCharter, "id = ?", id)
	if pCharter.ID == 0 {
		return pCharter, errors.New("id not found")
	}
	return pCharter, nil
}

// never gonna give you up
