package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID 			 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  	 string     `gorm:"unique;not null" json:"username"`
	Email 		 string     `gorm:"unique;not null" json:"email"`
	PasswordHash string     `gorm:"not null" json:"-"`
	CreatedAt 	 time.Time  `json:"created_at"`
	UpdatedAt 	 time.Time  `json:"updated_at"`
}