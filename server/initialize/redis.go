package initialize

import (
	"context"
	"fmt"
	"github.com/lindocedskes/global"
	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	redisCfg := global.NBUCTF_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	ok, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println("redis connect ping failed, err:", err)
		return nil
	} else {
		fmt.Println("redis connect ping response:", ok)
		return client
	}
}
