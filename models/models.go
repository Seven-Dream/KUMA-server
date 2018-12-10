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

	err = tableCheckAndCreate(db, &User{})
	if err != nil  {
		fmt.Println("User")
		panic(err)
	}

	err = tableCheckAndCreate(db, &Lecture{})
	if err != nil  {
		fmt.Println("Lecture")
		panic(err)
	}
	err = tableCheckAndCreate(db, &Week{})
	if err != nil  {
		fmt.Println("Week")
		panic(err)
	}
	err = tableCheckAndCreate(db, &Time{})
	if err != nil  {
		fmt.Println("Time")
		panic(err)
	}

	/*
	db, err := gorm.Open("postgres", "host=localhost user=tetsuya dbname=kuma sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(reflect.TypeOf(db))
	*/
}

// dbとの接続
func open() (*gorm.DB, error){
	return gorm.Open("postgres", "host=localhost user=tetsuya dbname=kuma sslmode=disable")
}

// テーブルが存在すれば作成する
func tableCheckAndCreate(db *gorm.DB, t interface{}) error {
	ok := db.HasTable(t)
	if !ok {
		err := db.CreateTable(t).Error
		return err
	}
	return nil
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
		fmt.Println(err)
	}

	return lecture
}
