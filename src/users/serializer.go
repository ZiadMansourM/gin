package users

type UserSerializer struct {
	Id    uint64 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
}
