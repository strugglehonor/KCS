package redis

import (
	redisconf "github.com/strugglehonor/KCS/internal/config/redis"
	"github.com/strugglehonor/KCS/pkg/log"
	"github.com/strugglehonor/KCS/pkg/redis"
)

func InitRedis() (*redis.RedisCli) {
	log.INFO.Println("init redis")
	option := redis.NewDefaultOption()
	option.URL = redisconf.URL
	option.Password = redisconf.Password
	redisCli := redis.NewRedisCli(option)
	return redisCli
}
