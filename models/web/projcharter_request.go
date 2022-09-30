package web

import (
	"time"
)

type ProjectCharterRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	User_id     uint64     `json:"user_id"`
	UpdatedBy   string     `json:"updated_by"`
	DeletedBy   string     `json:"deleted_by"`
}

type ProjectCharterUpdateRequest struct {
	ID          uint       `json:"id" binding:"required"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	User_id     uint64     `json:"user_id"`
	UpdatedBy   string     `json:"updated_by" binding:"required"`
	DeletedBy   string     `json:"deleted_by"`
}
