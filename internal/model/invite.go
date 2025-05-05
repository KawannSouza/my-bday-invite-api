package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invite struct {
	ID 			uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID 		uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title 		string 	  `json:"title"`
	Description string    `json:"description"`
	EventDate   time.Time `json:"event_date"`
	Background  string    `json:"background"`
	Code 	    string    `gorm:"uniqueIndex" json:"code"`
	CreatedAt   time.Time 
	UpdatedAt   time.Time
}

func (i *Invite) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	i.Code = uuid.New().String()[:8]
	return
}