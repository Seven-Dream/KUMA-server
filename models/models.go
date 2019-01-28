package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
)


func init() {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.CreateTable(&User{}).Error
	if err != nil {
		panic(err)
	}
	err = db.CreateTable(&Lecture{}).Error
	if err != nil {
		panic(err)
	}
	err = db.CreateTable(&Other{}).Error
	if err != nil {
		panic(err)
	}
	err = db.CreateTable(&UniversityEvent{}).Error
	if err != nil {
		panic(err)
	}
	err = db.CreateTable(&StudentEvent{}).Error
	if err != nil {
		panic(err)
	}
	err = db.CreateTable(&Test{}).Error
	if err != nil {
		panic(err)
	}
	err = db.CreateTable(&Cancel{}).Error
	if err != nil {
		panic(err)
	}
	err =  db.CreateTable(&ChangeRoom{}).Error
	if err != nil {
		panic(err)
	}

	err = addUserData("", "")
	if err != nil {
		panic(err)
	}
}

// dbとの接続
func open() (*gorm.DB, error){
	var connect string
	if password == "" {
		connect = fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable", hostname, user, dbname, port)
	} else {
		connect = fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", hostname, user, dbname, password, port)
	}
	return gorm.Open("postgres", connect)
}

func UserCheckFromIdAndPass(id string, pass string) (string, error) {
	db, err := open()
	defer db.Close()
	user := &User{Id: id, Password: pass}
	err = db.Where(user).First(user).Error
	if err != nil {
		return "", err
	}
	return id, err
}

func AllLecture() []Lecture {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	lecture := []Lecture{}

	err = db.Find(&lecture).Error
	if err != nil {
		fmt.Println("------------")
		fmt.Println(err)
		fmt.Println("------------")
	}

	// ここはgo routine使ったほうが早くなるはず
	for i, v := range lecture {
		o := []Other{}
		err = db.Model(&v).Related(&o).Error
		if err != nil {
			fmt.Println(err)
		}
		lecture[i].Others = o
	}
	fmt.Println(lecture)

	return lecture
}

func CreateLectureData(l *Lecture) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Create(l).Error
	if err != nil {
		fmt.Println("doesn't insert data")
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func DeleteLectureDataFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	deleteOther := Other{LectureID: id}
	deleteLecture := Lecture{Id: id}

	err = db.Delete(&deleteLecture).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = db.Debug().Delete(&deleteOther, "lecture_id = ?", id).Error

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetLectureDataFromId(id int) (Lecture, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	lecture := Lecture{Id: id}
	o := []Other{}

	err = db.Find(&lecture).Related(&o).Error
	if err != nil {
		return Lecture{}, err
	}

	lecture.Others = o

	return lecture, nil
}

func DeleteOtherFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Where("lecture_id = ?", id).Delete(&Other{LectureID: id}).Error
	return err
}

func UpdateLectureFromArgment(lecture *Lecture) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Model(lecture).Updates(lecture).Error
	return err
}
