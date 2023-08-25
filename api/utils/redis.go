package utils

import (
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() error {
	Password := EnvRedisPassword()
	Address := EnvRedisAddress()
	Master := EnvRedisMaster()
	redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    Master,
		SentinelAddrs: []string{Address},
		Password:      Password,
	})

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
func CloseRedis() error {
	if redisClient == nil {
		return nil
	}

	err := redisClient.Close()
	if err != nil {
		return err
	}

	return nil
}
