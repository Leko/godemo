package controller

import (
	"github.com/gin-gonic/gin"
	"godemo/database"
	"godemo/model"
	"net/http"
)

type todos struct{}

// TODO関連の処理を行うコントローラ
var Todos todos

// TODO一覧を返却する
func (u *todos) List(c *gin.Context) {
	var todos []model.Todo

	db := database.GetDB()
	user := c.MustGet("user").(model.User)

	q := db.Where("user_id = ?", user.ID)
	completed := c.DefaultQuery("completed", "")

	if completed != "" {
		q = q.Where("completed = ?", completed)
	}

	if err := q.Find(&todos).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	}
}

// TODOを登録する
func (u *todos) Create(c *gin.Context) {
	var todo model.Todo

	if c.BindJSON(&todo) != nil {
		return
	}

	db := database.GetDB()
	user := c.MustGet("user").(model.User)

	todo.UserID = user.ID
	if err := db.Save(&todo).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}
