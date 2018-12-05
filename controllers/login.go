package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"

	m "../models"

	"net/http"
)

func LoginGet(c *gin.Context) {
	// sessionチェック
	user, err := sessionCheck(c)
	// session情報あり
	if err == nil {
		c.Redirect(http.StatusFound, "/mypage")
		return
	}
	render(c, "login.tmpl", gin.H{"title": "Login Page"})
}

func LoginPost(c *gin.Context) {
	user, err := c.PostForm("id")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	plainTextPassword, err := c.PostForm("pass")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// 暗号化
	password := createEncryptedPassword(plainTextPassword)

	// ユーザ情報の確認
	id, err := m.UserCheckFromIdAndPass(user, password)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	render(c, "mypage.tmpl", gin.H{"title": "MyPage",}
}
