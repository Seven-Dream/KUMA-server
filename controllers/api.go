package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	m "KUMA-server/models"
)

func LectureAPI(c *gin.Context) {
	le := m.AllLecture()

	c.JSON(http.StatusOK, le)
}

func TestAPI(c *gin.Context) {
	test, _ := m.GetAllTest()
	c.JSON(http.StatusOK, test)
}

func CancelAPI(c *gin.Context) {
	cancel, _ := m.GetAllCancel()
	c.JSON(http.StatusOK, cancel)
}

func ChangeRoomAPI(c *gin.Context) {
	room, _ := m.GetAllChangeRoom()
	c.JSON(http.StatusOK, room)
}

func UniversityAPI(c *gin.Context) {
	university, _ := m.GetAllUniversityEvent()
	c.JSON(http.StatusOK, university)
}

func StudentAPI(c *gin.Context) {
	student, _ := m.GetAllStudentEvent()
	c.JSON(http.StatusOK, student)
}
