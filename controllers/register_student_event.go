package controllers

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"

	m "KUMA-server/models"
)

func RegisterStudentEventGet(c *gin.Context) {
	_, err := sessionCheckout(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	render(c, "edit_student_event.tmpl", gin.H{"title": "Register Student Event Page", "subtitle": "学生イベント新規登録", "event": m.StudentEvent{}, "url": "/student_event/register"})

}

func RegisterStudentEventPost(c *gin.Context) {
	_, err := sessionCheckout(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	event, ok := getStudentEventFromPost(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/student_event/register")
		return
	}

	err = m.CreateStudentEventFromArgment(&event)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/student_event/register")
		return
	}
	c.Redirect(http.StatusSeeOther, "/student_event")

}
