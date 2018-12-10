package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TopGet(c *gin.Context) {
	// session情報の取得
	id, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	render(c, "top.tmpl", gin.H{"title": "Top Page"})
}
