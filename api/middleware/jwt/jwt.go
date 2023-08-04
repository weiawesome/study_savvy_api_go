package jwt

import "study_savvy_api_go/internal/repository/redis"

type MiddlewareJwt struct {
	Repository *redis.Repository
}
