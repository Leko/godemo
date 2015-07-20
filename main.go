package main

import (
	"github.com/gin-gonic/gin"
	"godemo/controller"
	"godemo/database"
	"godemo/model"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	migrate()

	router := gin.Default()

	router.Static("/css", "./assets/dist/css")
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/", controller.Users.Top)
	router.GET("/login", controller.Users.Login)
	router.GET("/register", controller.Users.Register)
	router.POST("/authenticate", controller.Users.Authenticate)
	router.POST("/users/create", controller.Users.Create)

	router.Run(":" + port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}

func migrate() {
	db := database.GetDB()

	db.AutoMigrate(&model.User{})
}
