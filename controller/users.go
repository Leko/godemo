package controller

import (
	"github.com/gin-gonic/gin"
	"godemo/database"
	"godemo/model"
	"net/http"
)

var Users users = users{}

type users struct{}

func (u *users) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "user_form.tpl", gin.H{})
}

func (u *users) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "user_form.tpl", gin.H{
		"new": true,
	})
}

func (u *users) Create(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := model.User{
		Email:    email,
		Password: model.PasswordHash(password),
	}

	db := database.GetDB()
	db.Create(&user)

	c.Redirect(http.StatusMovedPermanently, "/login")
}
