package controllers

import (
	"github.com/gin-gonic/gin"

	m "KUMA-server/models"

	"net/http"
	"fmt"
)

func LoginGet(c *gin.Context) {
	// sessionチェック
	_, err := sessionCheck(c)
	// session情報あり
	if err == nil {
		c.Redirect(http.StatusFound, "/top")
		return
	}
	render(c, "login.tmpl", gin.H{"title": "Login Page"})
}

func LoginPost(c *gin.Context) {
	// postからidの情報を取得
	user := c.PostForm("id")
	if user == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// postからpassの情報を取得
	plainTextPassword := c.PostForm("pass")
	if plainTextPassword == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// 暗号化
	password := createEncryptedPassword(plainTextPassword)

	// ユーザ情報の確認
	_, err := m.UserCheckFromIdAndPass(user, password)

	// ユーザ情報がない場合
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// sessionの情報を追加
	err = sessionAdd(c, user, "session")

	if err != nil {
		fmt.Println("not register session")
	}

	c.Redirect(http.StatusFound, "/top")
}
