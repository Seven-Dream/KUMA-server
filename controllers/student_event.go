package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
	"fmt"

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

	render(c, "edit_student_event.tmpl", gin.H{"title": "Edit Student Event Page", "subtitle": "学生イベント変更・削除", "event": studentEvent, "url": "/student_event/change/"+idStr})

}

func ChangeStudentEvent(c *gin.Context) {
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

	existingStudentEvent, _ := m.GetStudentEventFromId(id)

	postedStudentEvent, ok := getStudentEventFromPost(c)

	if !ok {
		c.Redirect(http.StatusSeeOther, "/student_event")
		return
	}

	if existingStudentEvent.Url != "" {
		if postedStudentEvent.Url == "" {
			postedStudentEvent.Url = existingStudentEvent.Url
		} else {
			err = deleteFile(existingStudentEvent.Url)
			if err != nil {
				fmt.Println(err)
				fmt.Println("can't delete file")
			}
		}
	}

	err = m.UpdateStudentEventFromArgment(&postedStudentEvent)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/student_event")
		return
	}

	c.Redirect(http.StatusSeeOther, "/student_event")
}

func DeleteStudentEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/student_event")
		return
	}

	StudentEvent, err := m.GetStudentEventFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/student_event")
		return
	}

	willDeleteFile := StudentEvent.Url

	err = m.DeleteStudentEventFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/student_event")
		return
	}

	if willDeleteFile != "" {
		err = deleteFile(willDeleteFile)
		if err != nil {
			fmt.Println(err)
			fmt.Println("can't delete image file")
		}
	}

	c.Redirect(http.StatusSeeOther, "/student_event")
}
