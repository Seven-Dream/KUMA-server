package controllers

import (
	"github.com/gin-gonic/gin"
)

func TopGet(c *gin.Context) {
	// session情報の取得
	id, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	render(c, "top.tmpl", gin.H{"title": "Top Page"})
}
