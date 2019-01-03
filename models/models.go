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
