package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Balance   uint32 `gorm:"default:0"`
}

type CreateUser struct {
	Name     string
	Email    string
	Password string
}
