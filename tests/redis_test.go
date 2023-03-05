package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/database"
)

func Test_redis(t *testing.T) {
	my_redis := database.MyRedis(
		"127.0.0.1",
		"6379",
		"4",
	)

	//my_redis.Set("hi", "you", 3000)
	value, err := my_redis.Get("hi")
	if err != nil {
		println(value)
	}
}
