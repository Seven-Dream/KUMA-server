package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	"net/http"
	"encoding/hex"
	"crypto/sha256"
	"errors"
)

func render(c *gin.Context, fileName string, values interface{}) {
	c.HTML(http.StatusOK, fileName, values)
}

func sessionCheck(c *gin.Context) (string, error) {
	store := ginsession.FromContext(c)
	id, ok := store.Get("id")
	if !ok {
		return "", errors.New("not found")
	}
	_, ok = store.Get("secret")
	if !ok {
		return "", errors.New("not found")
	}
	return id.(string), nil
}

func sessionAdd(c *gin.Context, id string, session string) error {
	store := ginsession.FromContext(c)
	store.Set("id", id)
	store.Set("secret", session)
	err := store.Save()
	return err
}

func createEncryptedPassword(plainTextPassword string) string {
	b := []byte(plainTextPassword)
	hashbyte := sha256.Sum256(b)

	password := hex.EncodeToString(hashbyte[:])
	return password
}
