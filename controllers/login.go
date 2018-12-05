package controllers

import (
	"github.com/gin-gonic/gin"

	m "KUMA-server/models"

	"net/http"
)

func LoginGet(c *gin.Context) {
	// sessionチェック
	_, err := sessionCheck(c)
	// session情報あり
	if err == nil {
		c.Redirect(http.StatusFound, "/mypage")
		return
	}
	render(c, "login.tmpl", gin.H{"title": "Login Page"})
}

func LoginPost(c *gin.Context) {
	user := c.PostForm("id")
	if user == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	plainTextPassword := c.PostForm("pass")
	if plainTextPassword == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// 暗号化
	password := createEncryptedPassword(plainTextPassword)

	// ユーザ情報の確認
	_, err := m.UserCheckFromIdAndPass(user, password)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	render(c, "mypage.tmpl", gin.H{"title": "MyPage",})
}
