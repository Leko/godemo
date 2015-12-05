package session

import (
	"github.com/gorilla/sessions"
	"godemo/database"
	"godemo/model"
	. "gopkg.in/boj/redistore.v1"
	"net/http"
	"os"
	"strconv"
)

const keySession = "sessions"
const defaultSessionMaxAge = 30 * 24 * 3600

var (
	store *RediStore
)

func init() {
	var err error
	max := maxAge()

	store, err = NewRediStoreWithPool(database.GetRedisPool(), []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	store.SetMaxAge(max)
}

func GetSession(req *http.Request) *sessions.Session {
	session, err := store.Get(req, keySession)
	if err != nil {
		panic(err)
	}

	return session
}

func GetCurrentUser(req *http.Request) model.User {
	var user model.User

	id := GetSession(req).Values["userId"]

	if id != nil {
		db := database.GetDB()
		db.First(&user, id)
	} else {
		user = model.User{}
	}

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

func maxAge() int {
	env := os.Getenv("SESSION_MAX_AGE")
	if env == "" {
		return defaultSessionMaxAge
	}

	max, _ := strconv.Atoi(env)
	return max
}
