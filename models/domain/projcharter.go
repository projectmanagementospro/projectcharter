package domain

import (
	"time"

	"gorm.io/gorm"
)

type ProjectCharter struct {
	gorm.Model
	Name        string     `json:"name" gorm:"type:varchar(255);not null, unique"`
	Description string     `json:"description" gorm:"type:varchar(255);not null"`
	StartDate   *time.Time `json:"start_date"gorm:"type:timestamp"`
	EndDate     *time.Time `json:"end_date"gorm:"type:timestamp"`
	User_id     uint64     `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy   string     `json:"updated_by"`
	DeletedBy   string     `json:"deleted_by"`
}
