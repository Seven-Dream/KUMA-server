package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

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
