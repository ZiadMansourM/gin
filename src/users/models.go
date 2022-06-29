package users

import "time"

type UserModel struct {
	Id        uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt DeletedAt `gorm:"index"`
}

type Tabler interface {
	TableName() string
}

func (UserModel) TableName() string {
	return "users"
}
