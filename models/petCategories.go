package models

import (
	"time"
)

type PetCategories struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"not_null"`
	Status    uint      `json:"status"`
	CreatedBy uint      `json:"created_by" gorm:"default:1"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy uint      `json:"updated_by" gorm:"default:1"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
