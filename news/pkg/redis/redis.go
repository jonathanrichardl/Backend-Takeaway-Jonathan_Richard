package redis

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Redis   *redis.Client
	Context context.Context
	Timeout time.Duration
}

func NewRedisClient(Address string, Port int, Password string, DB int, Expiration int) *RedisClient {
	address := fmt.Sprintf("%s:%d", Address, Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: Password,
		DB:       DB,
	})
	context := context.Background()
	timeout := time.Duration(Expiration * int(time.Second))
	return &RedisClient{Redis: client, Context: context, Timeout: timeout}

}

func (R *RedisClient) SetData(Key string, Data string) error {
	err := R.Redis.Set(R.Context, Key, Data, R.Timeout)
	return err.Err()

}

func (R *RedisClient) GetData(Key string) (string, error) {
	data, err := R.Redis.Get(R.Context, Key).Result()
	return data, err
}
