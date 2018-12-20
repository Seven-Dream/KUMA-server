package models


type User struct {
	Id	string	`gorm:"type:varchar(64);primary_key;not null"`
	Password	string	`gorm:"type:char(64);not null"`
}

type Lecture struct {
	Id	int
	LectureName	string
	Teachar	string
	ClassRoom	string
	Year	int
	Quarter	int
	Others	[]Other
}

type Other struct {
	LectureID	int
	Week	string
	Time	int
}

