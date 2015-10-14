package database

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/joho/godotenv/autoload"
	"github.com/soveran/redisurl"
	"os"
	"strconv"
)

const defaultRedisMaxConnection = 30
const defaultRedisUrl = ":6379"

var (
	redisPool *redis.Pool
)

func init() {
	max := maxConnection()
	url := address()
	redisPool = redis.NewPool(func() (redis.Conn, error) {
		return redisurl.ConnectToURL(url)
	}, max)
}

func GetRedisPool() *redis.Pool {
	return redisPool
}

func maxConnection() int {
	env := os.Getenv("REDIS_MAX_CONNECTION")
	if env == "" {
		return defaultRedisMaxConnection
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
