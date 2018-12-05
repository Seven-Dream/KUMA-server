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


	router.Run(":8080")
}
