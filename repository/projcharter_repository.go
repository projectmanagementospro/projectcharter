package repository

import (
	"errors"
	"projectcharter/models/domain"

	"gorm.io/gorm"
)

type ProjectCharterRepository interface {
	All() []domain.ProjectCharter
	Create(projectcharter domain.ProjectCharter) domain.ProjectCharter
	Update(projectcharter domain.ProjectCharter) domain.ProjectCharter
	Delete(projectcharter domain.ProjectCharter)
	FindById(id uint) (domain.ProjectCharter, error)
}

type ProjectCharterConnection struct {
	//Connetion to database
	connection *gorm.DB
}

func NewProjectCharterRepository(db *gorm.DB) ProjectCharterRepository {
	return &ProjectCharterConnection{connection: db} //
}

func (conn *ProjectCharterConnection) All() []domain.ProjectCharter {
	var projectcharters []domain.ProjectCharter
	conn.connection.Find(&projectcharters)
	return projectcharters
}

func (conn *ProjectCharterConnection) Create(projectcharter domain.ProjectCharter) domain.ProjectCharter {
	conn.connection.Save(&projectcharter)
	return projectcharter
}

func (conn *ProjectCharterConnection) Update(projectcharter domain.ProjectCharter) domain.ProjectCharter {
	conn.connection.Omit("created_at").Save(&projectcharter)
	return projectcharter
}

func (conn *ProjectCharterConnection) Delete(projectcharter domain.ProjectCharter) {
	conn.connection.Delete(&projectcharter)
}

func (conn *ProjectCharterConnection) FindById(id uint) (domain.ProjectCharter, error) {
	var projectcharter domain.ProjectCharter
	conn.connection.Find(&projectcharter, "id = ?", id)
	if projectcharter.ID == 0 {
		return projectcharter, errors.New("id not found")
	}
	return projectcharter, nil
}
