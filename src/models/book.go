package models

type Book struct {
	Id       uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title    string `json:"title" binding:"required"`
	Author   User   `json:"author" gorm:"foreignkey:PersonID"`
	PersonID uint64 `json:"person_id"`
}
