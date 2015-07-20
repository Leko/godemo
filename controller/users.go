package controller

import (
	"github.com/gin-gonic/gin"
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
