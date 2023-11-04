package models

type User struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"name" form:"name"`
	Email string `gorm:"type:varchar(255)" json:"email" form:"email"`
	Password string `gorm:"type:varchar(255)" json:"password" form:"password"`
	Posts []Post
}