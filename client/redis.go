package client

import (
	"GoWeb-template/config"
	"context"
	"github.com/Echo7659/elog"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedisClient() {
	conf := config.Scd.Redis
	// 初始化连接
	conn := redis.NewClient(&redis.Options{
		Addr:     conf.GetDSN(),
		Password: conf.Password,
		DB:       conf.Db,
	})
	if err := conn.Ping(context.Background()).Err(); err != nil {
		elog.Say.Panicf("Redis连接初始化失败: %v", err)
	} else {
		elog.Say.Debug("Redis连接初始化成功")
	}
	Redis = conn
}
