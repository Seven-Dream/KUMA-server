package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"fmt"
	"strconv"
	"strings"

	m "KUMA-server/models"
)

func ShowUniversityEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	university, err := m.GetAllUniversityEvent()
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/top")
		return
	}
	render(c, "all_university_event.tmpl", gin.H{"title": "University Event Page", "university": university})

}

func RegisterUniversityEventGet(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	render(c, "edit_university_event.tmpl", gin.H{"title": "Register University Event Page", "subtitle": "大学予定新規登録", "university": m.UniversityEvent{}, "url": "/university_event/register"})
}

func RegisterUniversityEventPost(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	postedUniversity, ok := getUniversityEventFromPosted(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/university_event/register")
		return
	}

	err = m.CreateUniversityEventFromArgment(&postedUniversity)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/university_event/register")
		return
	}

	c.Redirect(http.StatusSeeOther, "/university_event")
}

func EditUniversityEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/university_event")
		return
	}

	editUniversity, err := m.GetUniversityEventFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/university_event")
		return
	}

	render(c, "edit_university_event.tmpl", gin.H{"title": "Edit University Event Page", "subtitle": "大学予定変更・削除", "university": editUniversity, "url": "/university_event/change/"+idStr})
}

func ChangeUniversityEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/university_event")
		return
	}

	updateUniversity, ok := getUniversityEventFromPosted(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/university_event/"+idStr)
		return
	}

	updateUniversity.Id = id
	err = m.UpdateUniversityEventFromArgment(&updateUniversity)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/university_event")
		return
	}
	c.Redirect(http.StatusSeeOther, "/university_event")
}

func  DeleteUniversityEvent(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/university_event")
		return
	}

	err = m.DeleteUniversityEventFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/university_event")
		return
	}
	c.Redirect(http.StatusSeeOther, "/university_event")
}

func getUniversityEventFromPosted(c *gin.Context) (m.UniversityEvent, bool) {
	allValueExist := true

	name    := c.PostForm("event_name")
	date    := c.PostForm("date"      )
	comment := c.PostForm("comment"   )

	if name == "" { allValueExist = false }
	if date == "" { allValueExist = false }
	// commentは空でも可能

	correctFormatDate := true
	splitedDate := strings.Split(date, "-")

	year, err := strconv.Atoi(splitedDate[0])
	if err != nil {
		correctFormatDate = false
	}
	month, err := strconv.Atoi(splitedDate[1])
	if err != nil {
		correctFormatDate = false
	}
	day, err := strconv.Atoi(splitedDate[2])
	if err != nil {
		correctFormatDate = false
	}

	returnUniversityEvent := m.UniversityEvent{}

	if correctFormatDate {
		returnUniversityEvent.Name    = name
		returnUniversityEvent.Year    = year
		returnUniversityEvent.Month   = month
		returnUniversityEvent.Day     = day
		returnUniversityEvent.Date    = date
		returnUniversityEvent.Comment = comment
	} else if allValueExist {
		returnUniversityEvent.Name = name
	}

	return returnUniversityEvent, (correctFormatDate && allValueExist) }
