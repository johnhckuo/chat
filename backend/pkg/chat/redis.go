package chat

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedis(connectionString string) Message {

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	client := redis.NewClient(opt)
	// 通過 cient.Ping() 來檢查是否成功連線到了 redis 伺服器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return &Redis{client: client, ctx: ctx}
}

func (r *Redis) Pop(key string) (*string, error) {
	res := r.client.RPop(key)

	if err := res.Err(); err != nil {
		return nil, err
	}

	resString := res.String()
	return &resString, nil
}

func (r *Redis) Push(key, value string) error {
	res := r.client.LPush(key, value)
	if err := res.Err(); err != nil {
		return err
	}
	return nil
}
