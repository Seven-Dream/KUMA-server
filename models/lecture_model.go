package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

func SearchLectureFromArgment(l *Lecture) ([]Lecture, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	o := l.Others[0]
	l.Others = nil

	returnLecture := []Lecture{}

	err = db.Debug().Select("DISTINCT id").Joins("JOIN others ON others.lecture_id = lectures.id").Where(l).Where(&o).Find(&returnLecture).Error
	if err != nil {
		return returnLecture, err
	}

	requireId := []int{}
	for _, v := range returnLecture {
		requireId = append(requireId, v.Id)
	}

	returnLecture, err = GetLectureFromMultiId(requireId)

	fmt.Println(returnLecture)

	return returnLecture, err
}

func GetLectureFromMultiId(id []int) ([]Lecture, error){
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	returnLecture := []Lecture{}

	err = db.Where("id IN (?)", id).Find(&returnLecture).Error
	if err != nil {
		return returnLecture, err
	}

	for i, v := range returnLecture {
		o := []Other{}
		err = db.Model(&v).Related(&o).Error
		if err != nil {
			fmt.Println(err)
		}
		returnLecture[i].Others = o
	}

	return returnLecture, err
}
