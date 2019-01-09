package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

	m "KUMA-server/models"
)

func ShowStudentEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	studentEvent, err := m.GetAllStudentEvent()
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/top")
	}

	render(c, "all_student_event.tmpl", gin.H{"title": "Student Event Page", "student": studentEvent})
}

func EditStudentEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/student_event")
		return
	}

	studentEvent, err := m.GetStudentEventFromId(id)

	render(c, "edit_student_event.tmpl", gin.H{"title": "Edit Student Event Page", "subtitle": "学生イベント変更・削除", "event": studentEvent, "url": "/student_event/change"})

}
