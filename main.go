package main

import (
	"github.com/gin-gonic/gin"
	"godemo/controller"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	router := gin.Default()

	router.Static("/css", "./assets/dist/css")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/login", controller.Users.Login)
	router.GET("/register", controller.Users.Register)

	router.Run(":" + port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}
