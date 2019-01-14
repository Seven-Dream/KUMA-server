package models


type User struct {
	Id       string `gorm:"type:varchar(64);primary_key;not null"`
	Password string `gorm:"type:char(64);not null"`
}

type Lecture struct {
	Id          int
	LectureName string
	Teachar     string
	ClassRoom   string
	Year        int
	Quarter     int
	Others      []Other
}

type Other struct {
	LectureID int
	Week      string
	Time      int
}

type UniversityEvent struct {
	Id      int
	Name    string
	Year    int
	Month   int
	Day     int
	Comment string `gorm"type:text"`
	Date    string
}

type StudentEvent struct {
	Id    int
	Name  string
	Year  int
	Month int
	Day   int
	Url   string
	Date  string
}

