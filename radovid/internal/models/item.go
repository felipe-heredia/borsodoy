package models

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Price       uint32    `json:"price"`
	ImageUrl    *string   `json:"image_url"`
	UserID      uuid.UUID `json:"user_id"`
	ExpiredAt   time.Time `json:"expired_at"`

	User *User `json:"user"`
  Bids *[]Bid `json:"bids"`
}

type CreateItem struct {
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Price       uint32    `json:"price"`
	ImageUrl    *string   `json:"image_url"`
	UserID      uuid.UUID `json:"user_id"`

	// Time in minutes
	ExpiresIn uint32 `json:"expires_in"`
}
