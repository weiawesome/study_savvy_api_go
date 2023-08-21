package redis

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis"
	"time"
)

const hashMapName = "mail_id_hash"

func GetHashValue(mail string, r *redis.Client) (string, error) {
	if mail == "system" {
		return mail, nil
	}
	ctx := context.Background()
	exists, err := r.Exists(ctx, hashMapName).Result()
	if err != nil {
		return "", err
	}

	if exists == 1 {
		value, err := r.HGet(ctx, hashMapName, mail).Result()
		if err != nil {
			return "", err
		}
		if value != "" {
			return value, nil
		} else {
			id := uuid.New().String()
			r.HSet(ctx, hashMapName, mail, id)
			return id, nil
		}
	} else if exists == 0 {
		r.HSet(ctx, hashMapName, "system", "system")
		expirationTime := time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour)
		err := r.ExpireAt(ctx, hashMapName, expirationTime).Err()
		if err != nil {
			return "", err
		}
		id := uuid.New().String()
		err = r.HSet(ctx, hashMapName, mail, id).Err()
		return id, err
	} else {
		return "", err
	}
}
