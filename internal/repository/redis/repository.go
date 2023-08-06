package redis

import (
	"github.com/redis/go-redis"
	"study_savvy_api_go/api/utils"
)

type Repository struct {
	client *redis.Client
}

func NewRepository() *Repository {
	return &Repository{client: utils.GetRedisClient()}
}
