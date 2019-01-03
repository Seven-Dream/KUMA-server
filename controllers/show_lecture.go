package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
	"fmt"

	m "KUMA-server/models"
)

func ShowLectureGet(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture")
		return
	}

	lecture, err := m.GetLectureDataFromId(id)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, "error")
	}
	render(c, "register_lecture.tmpl", gin.H{"title": "Lecture Change", "id": idStr, "lecture": lecture, "subtitle": "講義情報変更・削除", "url": "/lecture/change/"+idStr})
}
