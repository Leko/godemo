package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":" + port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}
