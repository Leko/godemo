package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"godemo/controller"
	"godemo/database"
	"godemo/model"
	"gopkg.in/bluesuncorp/validator.v5"
	"io"
	"net/http"
	"os"
	"unicode"
)

const defaultPort = "8080"

var (
	msgInvalidJSON     = "Invalid JSON format"
	msgInvalidJSONType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
	msgValidationFailed = func(e *validator.FieldError) string {
		switch e.Tag {
		case "required":
			return toSnakeCase(e.Field) + ": required"
		case "max":
			return toSnakeCase(e.Field) + ": too_long"
		default:
			return e.Error()
		}
	}
)

func main() {
	migrate()

	router := gin.Default()
	router.Static("/css", "./assets/dist/css")
	router.LoadHTMLGlob("templates/*")

	web := router.Group("", gin.WrapH(nosurf.NewPure(router)))
	{
		web.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		web.GET("/", controller.Users.Top)
		web.GET("/login", controller.Users.Login)
		web.GET("/logout", controller.Users.Logout)
		web.GET("/register", controller.Users.Register)
		web.POST("/authenticate", controller.Users.Authenticate)
		web.POST("/users/create", controller.Users.Create)
	}

	api := router.Group("/api")
	api.Use(apiHandle())
	{
		api.GET("/todos", controller.Todos.List)
		api.POST("/todos", controller.Todos.Create)
	}

	http.ListenAndServe(":"+port(), router)
}

func apiHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		db := database.GetDB()
		if err := db.Where("token = ?", c.Request.Header.Get("X-GODEMO-TOKEN")).Find(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"User not found"}})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()

		errs := make([]string, 0, len(c.Errors))
		for _, e := range c.Errors {
			// 1. エラーの種類で判定
			switch e.Err {
			case io.EOF:
				errs = append(errs, msgInvalidJSON)
			default:
				// 2. エラーの型で判定
				switch e.Err.(type) {
				case *json.SyntaxError:
					errs = append(errs, msgInvalidJSON)
				case *json.UnmarshalTypeError:
					errs = append(errs, msgInvalidJSONType(e.Err.(*json.UnmarshalTypeError)))
				case *validator.StructErrors:
					for _, fieldErr := range e.Err.(*validator.StructErrors).Flatten() {
						errs = append(errs, msgValidationFailed(fieldErr))
					}
				default:
					errs = append(errs, e.Error())
				}
			}
		}

		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{"errors": errs}) // -1 == not override the current error code
		}
	}
}

// https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
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

	db.AutoMigrate(&model.User{}, &model.Todo{})
}
