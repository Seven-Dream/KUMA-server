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
