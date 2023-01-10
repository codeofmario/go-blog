package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func InitRedis(settings *Settings) *redis.Client {
	options := &redis.Options{
		Addr:     settings.RedisAddr,
		Password: settings.RedisPassword,
		DB:       0,
	}
	client := redis.NewClient(options)
	checkRedis(client)

	return client
}

func checkRedis(client *redis.Client) {
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
