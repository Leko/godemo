package database

import (
	_ "github.com/joho/godotenv/autoload"
	. "gopkg.in/boj/redistore.v1"
	"os"
	"strconv"
)

const defaultRedisMaxConnection = 30
const defaultRedisMaxAge = 30 * 24 * 3600
const defaultRedisUrl = ":6379"

var redis *RediStore
var redisConnected bool

func GetKVS() *RediStore {
	if !redisConnected {
		redis = getConnection()
		redisConnected = true
	}

	return redis
}

func getConnection() *RediStore {
	max := maxConnection()
	url := address()
	// FIXME: Connection leek
	connection, err := NewRediStore(max, "tcp", url, "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	connection.SetMaxAge(maxAge())

	return connection
}

func maxConnection() int {
	env := os.Getenv("REDISCLOUD_MAX_CONNECTION")
	if env == "" {
		return defaultRedisMaxConnection
	}

	max, _ := strconv.Atoi(env)
	return max
}

func maxAge() int {
	env := os.Getenv("REDISCLOUD_MAX_AGE")
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
