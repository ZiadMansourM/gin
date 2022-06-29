package books

import "github.com/ziadmansourm/gin/users"

type BookModel struct {
	Id       uint64          `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title    string          `json:"title" binding:"required"`
	Author   users.UserModel `json:"author" gorm:"foreignkey:PersonID"`
	PersonID uint64          `json:"person_id"`
}

type Tabler interface {
	TableName() string
}

func (BookModel) TableName() string {
	return "books"
}
