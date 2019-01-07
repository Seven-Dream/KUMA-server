package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	"fmt"
	"net/http"
	"encoding/hex"
	"crypto/sha256"
	"errors"
	"strconv"
	"strings"
	"path/filepath"
	"os"

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

func getStudentEventFromPost(c *gin.Context) (m.StudentEvent, bool) {
	isAllPostDataExist := true

	// 指定しているkeyに対応した値がpostされいるか
	eventName, ok := isEmptyKey(c, "name")
	if !ok {
		isAllPostDataExist = false
	}
	date, ok := isEmptyKey(c, "date")
	if !ok {
		isAllPostDataExist = false
	}

	// 日付のを分割し、year -> month -> day
	splitedDate := strings.Split(date, "-")

	year, err := strconv.Atoi(splitedDate[0])
	if err != nil {
		isAllPostDataExist = false
	}
	month, err := strconv.Atoi(splitedDate[1])
	if err != nil {
		isAllPostDataExist = false
	}
	day, err := strconv.Atoi(splitedDate[2])
	if err != nil {
		isAllPostDataExist = false
	}

	// create StudentEvent struct
	studentEvent := m.StudentEvent{Name: eventName, Year: year, Month: month, Day: day, Date: date}

	if !isAllPostDataExist {
		return studentEvent, isAllPostDataExist
	}

	// pdfの確認
	pdf, err := c.FormFile("pdf")
	if err != nil {
		return studentEvent, isAllPostDataExist
	}

	// 拡張子の確認
	filename := filepath.Base(pdf.Filename)
	pos := strings.LastIndex(filename, ".")
	extension := filename[pos:]
	if extension != ".pdf" {
		return studentEvent, isAllPostDataExist
	}

	// pdfが存在するとサーバに作成
	saveFliePath := "./assets/img/" + filename
	fd, err := os.Create(saveFliePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fd.Close()
	}
	fmt.Println("---------")
	err = c.SaveUploadedFile(pdf, saveFliePath)
	if err != nil {
		fmt.Println(err)
		return studentEvent, isAllPostDataExist
	} else {
		studentEvent.Url = filename
		return studentEvent, isAllPostDataExist
	}
}
