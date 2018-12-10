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
	Weeks	[]Week
	Times	[]Time
}

type Week struct {
	LectureID	int
	Week	string
}

type Time struct {
	LectureID	int
	Time	int
}
