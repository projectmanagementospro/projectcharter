package repository

import (
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
