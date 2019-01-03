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

	err = db.DropTableIfExists(&User{}, &Lecture{}, &Other{}).Error
	if err != nil  {
		panic(err)
	}

	err = db.CreateTable(&User{}, &Lecture{}, &Other{}).Error
	if err != nil {
		panic(err)
	}

	err = db.DropTableIfExists(&StudentEvent{}).Error
	if err != nil  {
		panic(err)
	}

	err = db.CreateTable(&StudentEvent{}).Error
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
	return gorm.Open("postgres", "host=localhost user=tetsuya dbname=kuma sslmode=disable")
}

func UserCheckFromIdAndPass(id string, pass string) (string, error) {
	db, err := open()
	defer db.Close()
	user := &User{Id: id, Password: pass}
	err = db.First(user).Error
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

func GetLectureDataFromId(id int) (Lecture, error) {
  db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
  
  lecture := Lecture{Id: id}
	o := []Other{}

	err = db.First(&lecture).Related(&o).Error
	if err != nil {
		return Lecture{}, err
	}

	lecture.Others = o

	return lecture, nil
}

// 全ての学生イベント情報を取得
func GetAllStudentEvent() ([]StudentEvent, error){
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	event := []StudentEvent{}
	err = db.Find(&event).Error

	return event, err
}
