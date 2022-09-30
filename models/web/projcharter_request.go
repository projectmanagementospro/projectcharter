package web

import (
	"gorm.io/datatypes"
)

type ProjectCharterRequest struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	StartDate   datatypes.Date `json:"start_date"`
	EndDate     datatypes.Date `json:"end_date"`
	User_id     uint64         `json:"user_id"`
	UpdatedBy   string         `json:"updated_by"`
	DeletedBy   string         `json:"deleted_by"`
}

type ProjectCharterUpdateRequest struct {
	ID          uint           `json:"id" binding:"required"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	StartDate   datatypes.Date `json:"start_date"`
	EndDate     datatypes.Date `json:"end_date"`
	User_id     uint64         `json:"user_id"`
	UpdatedBy   string         `json:"updated_by" binding:"required"`
	DeletedBy   string         `json:"deleted_by"`
}
