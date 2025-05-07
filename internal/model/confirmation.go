package model

import (
	"time"

	"github.com/google/uuid"
)

type Confirmation struct {
	ID 		  uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	InviteID  uuid.UUID `gorm:"type:uuid;not null" json:"invite_id"`
	Name 	  string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

