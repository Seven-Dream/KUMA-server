package controllers

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	m "KUMA-server/models"
)

func RegisterLectureGet(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	o := []m.Other{}
	for range make([]int, 4) {
		o = append(o, m.Other{})
	}
	l := m.Lecture{Others: o}
	render(c, "register_lecture.tmpl", gin.H{"title": "Register Lecture Page", "subtitle": "新規登録", "url": "/lecture/register", "lecture": l})
}

func RegisterLecturePost(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	lecture, ok := getLectureDataFromPosted(c)
	if !ok {
		c.Redirect(http.StatusFound, "/lecture/register")
		return
	}
	fmt.Println(lecture)
	err = m.CreateLectureData(&lecture)
	if err != nil {
		c.Redirect(http.StatusFound, "/lecture/register")
	} else {
		c.Redirect(http.StatusFound, "/lecture")
	}
}
