package utils

import (
	"fmt"
	"github.com/redis/go-redis"
	"os"
)

var redisClient *redis.Client

func InitRedis() error {
	Password := os.Getenv("REDIS_PASSWORD")
	Port := os.Getenv("REDIS_PORT")
	Db := os.Getenv("REDIS_DB")
	opt, err := redis.ParseURL("redis://:" + Password + "@localhost:" + Port + "/" + Db)
	if err != nil {
		fmt.Println("AAA")
		return err
	}
	redisClient = redis.NewClient(opt)
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
