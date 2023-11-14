package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int `gorm:"primarykey" json:"id" yaml:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
