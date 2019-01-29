package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	m "KUMA-server/models"
)

func DeleteLecturePost(c *gin.Context) {
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
	err = m.DeleteLectureDataFromId(id)

	c.Redirect(http.StatusSeeOther, "/lecture")
}
