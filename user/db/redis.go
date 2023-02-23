package db

import (
	"fmt"
	"gopkg.in/redis.v4"
	"time"
)

var Client *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.33.9.239:8001",
		Password: "123#Redis",
		DB:       0,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败")
	}
	fmt.Println(pong)
	Client = client
}

func GetRedis() *redis.Client {
	return Client
}

func RedisSetString(key string, value string, time time.Duration) {
	//client := GetRedis()
	err := Client.Set(key, value, time).Err()
	if err != nil {
		fmt.Println(err)
	}
}
