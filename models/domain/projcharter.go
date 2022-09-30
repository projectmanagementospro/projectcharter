package domain

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProjectCharter struct {
	gorm.Model
	Name        string         `json:"name" gorm:"type:varchar(255);not null, unique"`
	Description string         `json:"description" gorm:"type:varchar(255);not null"`
	StartDate   datatypes.Date `json:"start_date"`
	EndDate     datatypes.Date `json:"end_date"`
	User_id     uint64         `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy   string         `json:"updated_by"`
	DeletedBy   string         `json:"deleted_by"`
}
