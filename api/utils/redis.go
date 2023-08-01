package utils

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var redisClient *redis.Client

func InitRedis() error {
	Password := os.Getenv("REDIS_PASSWORD")
	Port := os.Getenv("REDIS_PORT")
	Db := os.Getenv("REDIS_DB")
	opt, err := redis.ParseURL("redis://:" + Password + "@localhost:" + Port + "/" + Db)
	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(opt)

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
