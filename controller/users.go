package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"godemo/database"
	"godemo/model"
	"godemo/session"
	"net/http"
)

var Users users = users{}

type users struct{}

func (u *users) Top(c *gin.Context) {
	user := session.GetCurrentUser(c.Request)

	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"user": user,
	})
}

func (u *users) Login(c *gin.Context) {
	csrfToken := nosurf.Token(c.Request)
	c.HTML(http.StatusOK, "user_form.tpl", gin.H{
		"csrfToken": csrfToken,
	})
}

func (u *users) Register(c *gin.Context) {
	csrfToken := nosurf.Token(c.Request)
	c.HTML(http.StatusOK, "user_form.tpl", gin.H{
		"new":       true,
		"csrfToken": csrfToken,
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

func (u *users) Authenticate(c *gin.Context) {
	user := model.User{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	id, err := user.Auth()
	if err == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	s := session.GetSession(c.Request)
	s.Values["userId"] = id
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func (u *users) Logout(c *gin.Context) {
	session.Destroy(c.Request, c.Writer)
	c.Redirect(http.StatusMovedPermanently, "/")
}
