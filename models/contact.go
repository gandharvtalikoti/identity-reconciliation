package models

import (
	"time"
)

type Contact struct {
	ID             uint       `gorm:"primaryKey"`
	PhoneNumber    string     `gorm:"column:phone_number"`
	Email          string     `gorm:"column:email"`
	LinkedID       *uint      `gorm:"column:linked_id"` // Nullable
	LinkPrecedence string     `gorm:"column:link_precedence;type:varchar(10);check:link_precedence IN ('primary','secondary')"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"` // Nullable
}
