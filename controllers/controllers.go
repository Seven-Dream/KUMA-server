package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	"net/http"
	"encoding/hex"
	"crypto/sha256"
)

func render(c *gin.Context, fileName string, values interface{}) {
	c.HTML(http.StatusOK, fileName, values)
}

func sessionCheck(c *gin.Context) {
	store := ginsession.FromContext(c)
	id, err := store.Get("id")
	if err != nil {
		return _, err
	}
	secret, err := store.Get("secret")
	if err != nil {
		return _, err
	}
	return id, nil
}

func createEncryptedPassword(plainTextPassword string) string {
	b := []byte(plainTextPassword)
	hashbyte := sha256.Sum256(b)

	password := hex.EncodeToString(hashbyte[:])
	return password
}
