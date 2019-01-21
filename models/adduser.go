package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"encoding/hex"
	"crypto/sha256"

)

func createEncryptedPassword(plainTextPassword string) string {
	b := []byte(plainTextPassword)
	hashbyte := sha256.Sum256(b)

	password := hex.EncodeToString(hashbyte[:])
	return password
}

func addUserData(id string, pass string) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := "kuma"
	pass = "kuma"
	password := createEncryptedPassword(pass)

	err = db.Create(&User{Id: user, Password: password}).Error
	if err != nil {
		panic(err)
	}
	return err
}

