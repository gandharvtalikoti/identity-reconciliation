package models

import (
	"time"
)

type Contact struct {
	ID             uint       `gorm:"primaryKey"`
	PhoneNumber    *string    `gorm:"uniqueIndex"`
	Email          *string    `gorm:"uniqueIndex"`
	LinkedID       *uint
	LinkPrecedence string     `gorm:"type:enum('primary','secondary')"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
}
