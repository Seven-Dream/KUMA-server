package models


type User struct {
	Id       string `gorm:"type:varchar(64);primary_key;not null"`
	Password string `gorm:"type:char(64);not null"`
}

type Lecture struct {
	Id          int     `json:"lecture_id" gorm:"not null"`
	LectureName string  `json:"lecture_name" gorm:"not null"`
	Teachar     string  `json:"teacher"`
	ClassRoom   string  `json:"classroom"`
	Year        int     `json:"year"`
	Quarter     int     `json:"quarter"`
	Others      []Other `json:"week"`
}

type Other struct {
	LectureID  int    `json:"lecture_id"`
	Week       int    `json:"week"`
	Time       int    `json:"period"`
	WeekString string
}

type Test struct {
	LectureID int    `json:"lecture_id"`
	Id        int
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	ClassRoom string `json:"classroom"`
	Comment   string `json:"comment" gorm:"type:text"`
}

type Cancel struct {
	LectureID int    `json:"lecture_id"`
	Id        int
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	Comment   string `json:"comment" gorm:"type:text"`
}

type ChangeRoom struct {
	LectureID int    `json:"lecture_id"`
	Id        int
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	ClassRoom string `json:"classroom"`
}

type UniversityEvent struct {
	Id      int    `json:"id"`
	Name    string `json:"event_name"`
	Year    int    `json:"year"`
	Month   int    `json:"month"`
	Day     int    `json:"day"`
	Comment string `json:"comment" gorm"type:text"`
	Date    string
}

type StudentEvent struct {
	Id    int    `json:"id"`
	Name  string `json:"event_name"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
	Url   string `json:"url"`
	Date  string
}

