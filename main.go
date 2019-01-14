package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	c "KUMA-server/controllers"
)

func main() {
	router := gin.Default()

	router.Use(ginsession.New())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")

	router.GET("/login", c.LoginGet)
	router.POST("/login", c.LoginPost)

	router.GET("/lecture", c.LectureGet)

	router.GET("/lecture/register", c.RegisterLectureGet)
	router.POST("/lecture/register", c.RegisterLecturePost)

	router.POST("/lecture/delete/:id", c.DeleteLecturePost)

	router.GET("/lecture/show/:id", c.ShowLectureGet)

	router.GET ("/university_event",            c.ShowUniversityEvent)
	router.GET ("/university_event/register",   c.RegisterUniversityEventGet)
	router.POST("/university_event/register",   c.RegisterUniversityEventPost)
	router.GET ("/university_event/show/:id",   c.EditUniversityEvent)
	router.POST("/university_event/change/:id", c.ChangeUniversityEvent)
	router.POST("/university_event/delete/:id", c.DeleteUniversityEvent)

	router.GET("/student_event", c.ShowStudentEvent)
	router.GET("/student_event/register", c.RegisterStudentEventGet)
	router.POST("/student_event/register", c.RegisterStudentEventPost)
	router.GET("/student_event/show/:id", c.EditStudentEvent)
	router.POST("/student_event/change/:id", c.ChangeStudentEvent)
	router.POST("/student_event/delete/:id", c.DeleteStudentEvent)

	router.GET("/logout", c.LogoutGet)

	router.GET("/top", c.TopGet)

	router.Run(":8080")
}
