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

	router.GET("/student_event", c.ShowStudentEvent)

	router.GET("/logout", c.LogoutGet)

	router.GET("/top", c.TopGet)

	router.Run(":8080")
}
