package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
  ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
	Name      string     `json:"name" gorm:"name"`
  Email     string     `jgon:"email" gorm:"uniqueIndex"`
	Password  string     `json:"password" gorm:"password"`
  Balance   uint32     `json:"balance" gorm:"default:0"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
