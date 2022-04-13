package helpers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func init() {
	initRedis()
}

var ctx = context.Background()

var rdb *redis.Client

func Set(key string, value interface{}) error {
	err := rdb.Set(ctx, key, value, 20*time.Minute).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func Rpush(key string, value interface{}) error {
	err := rdb.RPush(ctx, key, value).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func Lpush(key string, value interface{}) error {

	err := rdb.LPush(ctx, key, value).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func Lrange(key string, start, stop int64) ([]string, error) {
	res, err := rdb.LRange(ctx, key, start, stop).Result()

	if err != nil {
		return nil, err
	}
	return res, nil
}

func Del(key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func GetCache() *redis.Client {
	return rdb
}

func Get(key string) (string, error) {

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
		return "", err
	}
	return value, nil
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
