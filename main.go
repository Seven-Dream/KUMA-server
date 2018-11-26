package main

import (
	"github.com/gin-gonic/gin"
	c "KUMA-server/controllers"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")

	router.GET("/login", c.Login)

	router.Run(":8080")
}
