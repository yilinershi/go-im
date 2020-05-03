package db

import (
	"fmt"
	"gopkg.in/redis.v4"
)

var RedisClient=new(redis.Client)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("server link redis err,err=", err)
		return
	}
	fmt.Println("server link redis success")
	fmt.Println(pong, err)
}
