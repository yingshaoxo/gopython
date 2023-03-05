// We use this package to handle databases
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yingshaoxo/gopython/string_tool"
)

type MyRedis_Type struct {
	Redis_client    redis.Client
	Context         context.Context
	cancel_function context.CancelFunc
}

func MyRedis(host string, port string, database_number string) MyRedis_Type {
	address := fmt.Sprintf("%s:%s", host, port)

	the_context, the_cancel_function := context.WithCancel(context.Background())

	raw_redis := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       string_tool.String_to_int(database_number, 0),
	})

	my_redis := MyRedis_Type{
		Redis_client:    *raw_redis,
		Context:         the_context,
		cancel_function: the_cancel_function,
	}

	return my_redis
}

func (self *MyRedis_Type) Stop() {
	defer self.cancel_function()
	if err := self.Redis_client.Close(); err != nil {
		panic(err)
	}
}

func (self *MyRedis_Type) Get(key string) (string, error) {
	value, err := self.Redis_client.Get(self.Context, key).Result()
	return value, err
}

func (self *MyRedis_Type) Set(key string, value string, timeout_in_milliseconds int) error {
	err := self.Redis_client.Set(self.Context, key, value, time.Duration(timeout_in_milliseconds)*time.Millisecond).Err()
	return err
}
