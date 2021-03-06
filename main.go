package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	c "KUMA-server/controllers"

	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.Use(ginsession.New())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")
	os.Mkdir("./assets/img/", 0666)


	router.GET("/", func(co *gin.Context) {
		co.Redirect(http.StatusSeeOther, "/login")
		return
	})
	router.GET("/login", c.LoginGet)
	router.POST("/login", c.LoginPost)

	router.GET("/lecture", c.LectureGet)

	router.GET("/lecture/register", c.RegisterLectureGet)
	router.POST("/lecture/register", c.RegisterLecturePost)
	router.POST("/lecture/delete/:id", c.DeleteLecturePost)
	router.GET("/lecture/show/:id", c.ShowLectureGet)
	router.POST("/lecture/change/:id", c.ChangeLecturePost)

	router.GET ("/lecture/search",                      c.SearchLecture)
	router.POST("/lecture/search/result",               c.SearchResultLecture)
	router.GET ("/lecture/option/list/:id",                  c.ShowLectureOption)
	router.GET ("/lecture/option/cancel/register/:lectureId",      c.RegisterCancelGet)
	router.GET ("/lecture/option/test/register/:lectureId",        c.RegisterTestGet)
	router.GET ("/lecture/option/change_room/register/:lectureId", c.RegisterChangeRoomGet)
	router.POST("/lecture/option/cancel/register",        c.RegisterCancelPost)
	router.POST("/lecture/option/test/register",          c.RegisterTestPost)
	router.POST("/lecture/option/change_room/register",   c.RegisterChangeRoomPost)
	router.GET ("/lecture/option/cancel/show/:id",        c.ShowCancel)
	router.GET ("/lecture/option/test/show/:id",          c.ShowTest)
	router.GET ("/lecture/option/change_room/show/:id",   c.ShowChangeRoom)
	router.POST("/lecture/option/cancel/delete/:id",      c.DeleteCancel)
	router.POST("/lecture/option/test/delete/:id",        c.DeleteTest)
	router.POST("/lecture/option/change_room/delete/:id", c.DeleteChangeRoom)

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

	router.GET("/api/lecture",     c.LectureAPI)
	router.GET("/api/test",        c.TestAPI)
	router.GET("/api/cancel",      c.CancelAPI)
	router.GET("/api/change_room", c.ChangeRoomAPI)
	router.GET("/api/university",  c.UniversityAPI)
	router.GET("/api/student",     c.StudentAPI)

	router.Run(":8080")
}
