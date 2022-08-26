package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
