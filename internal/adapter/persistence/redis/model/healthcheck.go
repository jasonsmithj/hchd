package model

import (
	"github.com/go-redis/redis/v7"
	"github.com/jasonsmithj/hchd/internal/adapter/configuration"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func redisConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configuration.GetRedisAddr(),
		Password: configuration.Get().RedisPassword,
		DB:       configuration.Get().RedisDB,
	})
}

func RedisHealthcheck() error {
	c := redisConnect()
	_, err := c.Ping().Result()

	defer c.Close()

	return err
}
