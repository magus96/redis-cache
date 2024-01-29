package main

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{client: c}
}

func (c *RedisCache) Get(key int) (string, bool) {
	ctx := context.Background()
	val, err := c.client.Get(ctx, strconv.Itoa(key)).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (c *RedisCache) Set(key int, val string) error {
	ctx := context.Background()
	_, err := c.client.Set(ctx, strconv.Itoa(key), val, 0).Result()
	return err
}

func (c *RedisCache) Remove(key int) error{
	ctx := context.Background()
	_, err := c.client.Del(ctx, strconv.Itoa(key)).Result()
	return err
}