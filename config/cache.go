package config

import (
	"context"
	"encoding/json"
	"time"
)

func CacheSet(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RedisClient.Set(context.Background(), key, data, ttl).Err()
}

func CacheGet(key string, dest interface{}) error {
	val, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return err // includes redis.Nil when key doesn't exist
	}
	return json.Unmarshal([]byte(val), dest)
}

func CacheDelete(key string) error {
	return RedisClient.Del(context.Background(), key).Err()
}