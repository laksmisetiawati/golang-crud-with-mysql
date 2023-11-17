package models

import (
	"time"
)

type Accounts struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Status    uint      `json:"status"`
	CreatedBy uint      `json:"created_by" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy uint      `json:"updated_by" gorm:"default:1"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
