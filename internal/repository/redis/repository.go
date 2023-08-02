package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis"
	"os"
	"strconv"
	"study_savvy_api_go/api/utils"
	"time"
)

type Repository struct {
	client *redis.Client
}

func NewRepository() *Repository {
	return &Repository{client: utils.GetRedisClient()}
}

func (r *Repository) SetToBlacklist(jwt string) error {
	ctx := context.Background()

	expiredDays, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_DAYS"))

	if expiredDays == 0 {
		expiredDays = 1
	}
	expireDuration := time.Duration(expiredDays) * 24 * time.Hour
	_, err := r.client.SAdd(ctx, jwt, "black_list", expireDuration).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ValidateInBlacklist(jwt string) error {
	ctx := context.Background()

	if exists, err := r.client.Exists(ctx, jwt).Result(); err != nil {
		return err
	} else if exists == 1 {
		return errors.New("jwt in black_list")
	} else {
		return nil
	}
}