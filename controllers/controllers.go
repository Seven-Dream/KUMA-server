package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	"net/http"
	"encoding/hex"
	"crypto/sha256"
	"errors"
	"strconv"

	m "KUMA-server/models"
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

func sessionDestroy(c *gin.Context) {
	ginsession.Destroy(c)
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

func getLectureDataFromPosted(c *gin.Context) (m.Lecture, bool) {
	l := m.Lecture{}
	noneFlag := false

	name, ok := isEmptyKey(c, "name")
	if !ok {
		noneFlag = true
	}
	teachar, ok := isEmptyKey(c, "teachar")
	if !ok {
		noneFlag = true
	}
	yearString, _ := isEmptyKey(c, "year")
	year, err := strconv.Atoi(yearString)
	if err != nil{
		noneFlag = true
	}
	room, ok := isEmptyKey(c, "room")
	if !ok {
		noneFlag = true
	}
	quarterString, _ := isEmptyKey(c, "quarter")
	quarter, err := strconv.Atoi(quarterString)
	if err != nil {
		noneFlag = true
	}

	other := c.PostFormArray("other[]")
	otherEmptyFlag := true
	o := []m.Other{}

	for _, v := range other {
		weekKey := "week"+v
		timeKey := "time"+v

		week, okw := isEmptyKey(c, weekKey)
		timeStr, okt := isEmptyKey(c, timeKey)
		time, err := strconv.Atoi(timeStr)
		if okw && okt && err == nil {
			otherEmptyFlag = false
			o = append(o, m.Other{Week: week, Time: time})
		}
	}
	if noneFlag || otherEmptyFlag {
		return m.Lecture{}, false
	}
	l = m.Lecture{LectureName: name, Teachar: teachar, ClassRoom: room, Year: year, Quarter: quarter, Others: o}
	return l, true
}

func isEmptyKey(c *gin.Context, key string) (string, bool) {
	value := c.PostForm(key)
	if value == "" {
		return "", false
	} else {
		return value, true
	}
}
