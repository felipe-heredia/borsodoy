package models

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`

	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      uint32    `json:"amount"`
	ItemID      uuid.UUID `json:"item_id"`
	UserID      uuid.UUID `json:"user_id"`

	User *User `json:"user"`
	Item *Item `json:"item"`
}

type CreateBid struct {
	Amount      uint32    `json:"amount"`
	WithdrawnIn uint32    `json:"withdrawn_in"`
	ItemID      uuid.UUID `json:"item_id"`
	UserID      uuid.UUID `json:"user_id"`
}
