package repository

import (
	"errors"
	"projectcharter/models/domain"

	"gorm.io/gorm"
)

type ProjectCharterRepository interface {
	All() []domain.ProjectCharter
	Create(p domain.ProjectCharter) domain.ProjectCharter
	Update(p domain.ProjectCharter) domain.ProjectCharter
	Delete(p domain.ProjectCharter)
	FindById(id uint64) (domain.ProjectCharter, error)
}

type ProjectCharterConnection struct {
	connection *gorm.DB
}

func NewProjectCharterRepository(connection *gorm.DB) ProjectCharterRepository {
	return &ProjectCharterConnection{connection: connection}
}

func (c *ProjectCharterConnection) All() []domain.ProjectCharter {
	var projectcharters []domain.ProjectCharter
	c.connection.Find(&projectcharters)
	return projectcharters
}

func (c *ProjectCharterConnection) Create(p domain.ProjectCharter) domain.ProjectCharter {
	c.connection.Save(&p)
	return p
}

func (c *ProjectCharterConnection) Update(p domain.ProjectCharter) domain.ProjectCharter {
	c.connection.Save(&p)
	return p
}

func (c *ProjectCharterConnection) Delete(p domain.ProjectCharter) {
	c.connection.Save(&p)
}

func (c *ProjectCharterConnection) FindById(id uint64) (domain.ProjectCharter, error) {
	var projectcharter domain.ProjectCharter
	c.connection.Find(&projectcharter, "id = ?", id)
	if projectcharter.ID == 0 {
		return projectcharter, errors.New("id not found")
	}
	return projectcharter, nil
}
