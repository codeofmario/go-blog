package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenRepository interface {
	GetToken(key string) (string, error)
	SaveToken(key string, token string, exp time.Time) (bool, error)
	DeleteToken(key string) (bool, error)
}

type TokenRepositoryImpl struct {
	Redis *redis.Client
}

func NewTokenRepository(redis *redis.Client) TokenRepository {
	return &TokenRepositoryImpl{Redis: redis}
}

func (r TokenRepositoryImpl) GetToken(key string) (string, error) {
	token, err := r.Redis.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	
	return token, nil
}

func (r TokenRepositoryImpl) SaveToken(key string, token string, exp time.Time) (bool, error) {
	end := time.Unix(exp.Unix(), 0)
	now := time.Now()
	token, err := r.Redis.Set(context.Background(), key, token, end.Sub(now)).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r TokenRepositoryImpl) DeleteToken(key string) (bool, error) {
	_, err := r.Redis.Del(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}
