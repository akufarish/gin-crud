package models

type Post struct {
	Id int64 `gorm:"primaryKey"`
	UserId uint
	Judul string `gorm:"type:varchar(255)"`
	Slug string `gorm:"type:varchar(255)"`
	Konten string `gorm:"type:varchar(255)"`
}