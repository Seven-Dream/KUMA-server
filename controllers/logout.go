package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	sessionDestroy(c)
	c.Redirect(http.StatusFound, "/login")
}
