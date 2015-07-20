package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"godemo/database"
	"godemo/model"
)

const keySession = "sessions"

func GetSession(req *http.Request) *sessions.Session {
	store := database.GetKVS()
	session, err := store.Get(req, keySession)
	if err != nil {
		panic(err)
	}

	return session
}

func GetCurrentUser(req *http.Request) model.User {
	var user model.User

	id := GetSession(req).Values["userId"]
	db := database.GetDB()
	db.First(&user, id)

	return user
}

func Save(req *http.Request, res http.ResponseWriter) {
	if err := sessions.Save(req, res); err != nil {
		panic(err)
	}
}

func Destroy(req *http.Request, res http.ResponseWriter) {
	session := GetSession(req)
	session.Options.MaxAge = -1
	Save(req, res)
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := GetSession(c.Request)

		if s.Values["userId"] == nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			c.Next()
		}
	}
}
