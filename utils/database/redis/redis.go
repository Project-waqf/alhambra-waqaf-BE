package redis

import (
	"wakaf/config"
	"wakaf/pkg/helper"

	"github.com/go-redis/redis/v8"
)

var log = helper.Logger()

func InitRedis(config *config.AppConfig) redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST + ":" + config.REDIS_PORT,
		Password: "",
		DB:       0,
	})
	return *redis
}
