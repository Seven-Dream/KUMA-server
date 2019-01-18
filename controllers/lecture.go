package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	m "KUMA-server/models"
)

func LectureGet(c *gin.Context) {
	// session確認
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// 講義情報を全て取得
	lecture := m.AllLecture()

	render(c, "lecture.tmpl", gin.H{"title": "Lecture Page", "lecture": lecture})
}

func fromNumberToWeek(relationToWeek int) string {
	var week string
	switch relationToWeek {
	case 1:
		week = "月曜日"
	case 2:
		week = "火曜日"
	case 3:
		week = "水曜日"
	case 4:
		week = "木曜日"
	case 5:
		week = "金曜日"
	case 6:
		week = "土曜日"
	}
	return week
}
