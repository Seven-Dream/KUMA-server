package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func render(c *gin.Context, fileName string, values interface{}) {
	c.HTML(http.StatusOK, fileName, values)
}

func Login(c *gin.Context) {
	render(c, "login.tmpl", gin.H{"title": "Login Page"}
}
