package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	c "./controllers"
)

func main() {
	router := gin.Default()

	router.Use(ginsession.New())

	router.GET("/login", c.LoginGet)
	router.POST("/login", c.LooginCheck(), c.LoginPost)

	router.Run(":8080")
}
