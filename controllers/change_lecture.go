package controllers

import (
	"github.com/gin-gonic/gin"

	"strconv"
	"net/http"
	"fmt"

	m "KUMA-server/models"
)

func ChangeLecturePost(c *gin.Context) {
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

	lecture, ok := getLectureDataFromPosted(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/lecture/show/"+idStr)
		return
	}

	err = m.DeleteOtherFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/show/"+idStr)
		return
	}

	lecture.Id = id

	err = m.UpdateLectureFromArgment(&lecture)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/show/"+idStr)
		return
	}
	c.Redirect(http.StatusSeeOther, "/lecture/show/"+idStr)
}
