package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutGet(c *gin.Context) {
	sessionDestroy(c)
	c.Redirect(http.StatusFound, "/login")
}
