package interfaces

import (
	"study_savvy_api_go/internal/repository/redis"
)

func RedisHashInterface(mail string) (string, error) {
	redisRepository := redis.NewRepository()
	return redisRepository.GetHashValue(mail)
}
