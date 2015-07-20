package database

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/joho/godotenv/autoload"
	"github.com/soveran/redisurl"
	. "gopkg.in/boj/redistore.v1"
	"os"
	"strconv"
)

const defaultRedisMaxConnection = 30
const defaultRedisMaxAge = 30 * 24 * 3600
const defaultRedisUrl = ":6379"

var conn *RediStore
var redisConnected bool

func GetKVS() *RediStore {
	if !redisConnected {
		conn = getConnection()
		redisConnected = true
	}

	return conn
}

func getConnection() *RediStore {
	max := maxConnection()
	url := address()
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redisurl.ConnectToURL(url)
	}, max)

	// FIXME: Connection leek
	connection, err := NewRediStoreWithPool(pool, []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	connection.SetMaxAge(maxAge())

	return connection
}

func maxConnection() int {
	env := os.Getenv("REDIS_MAX_CONNECTION")
	if env == "" {
		return defaultRedisMaxConnection
	}

	max, _ := strconv.Atoi(env)
	return max
}

func maxAge() int {
	env := os.Getenv("REDIS_MAX_AGE")
	if env == "" {
		return defaultRedisMaxAge
	}

	max, _ := strconv.Atoi(env)
	return max
}

func address() string {
	address := os.Getenv("REDISCLOUD_URL")
	if address == "" {
		return defaultRedisUrl
	}

	return address
}
