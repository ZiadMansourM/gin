package models

import (
	"time"
)

type User struct {
	Id        uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt DeletedAt `gorm:"index"`
}
