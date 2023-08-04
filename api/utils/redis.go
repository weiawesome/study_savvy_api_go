package utils

import (
	"github.com/redis/go-redis"
)

var redisClient *redis.Client

func InitRedis() error {
	Password := EnvRedisPassword()
	Address := EnvRedisAddress()
	Db := EnvRedisDb()
	opt, err := redis.ParseURL("redis://:" + Password + "@" + Address + "/" + Db)
	if err != nil {
		return err
	}
	redisClient = redis.NewClient(opt)
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
