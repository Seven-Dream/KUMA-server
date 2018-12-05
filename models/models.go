package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"reflect"
)

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=tetsuya dbname=kuma sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(reflect.TypeOf(db))
}

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
