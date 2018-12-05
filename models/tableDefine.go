package models


type User struct {
	Id	string	`gorm:"type:varchar(64);primary_key;not null"`
	Password	string	`gorm:"type:char(64);not null"`
}
