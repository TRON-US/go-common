package redis

import (
	"fmt"
	redis "github.com/go-redis/redis/v7"
)

type TGRDDB struct {
	rd *redis.Client
}

var (
	host = ""
	port = ""
	db   = 0
)

func ParseRedisURL(url string) (*redis.Options, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return opts, nil
}

func NewRedisConn(redisOptions *redis.Options) *TGRDDB {
	var client = redis.NewClient(redisOptions)
	return &TGRDDB{client}
}

func CheckRedisConnection(connection *TGRDDB) error {
	val, err := connection.rd.Ping().Result()
	if err != nil {
		return err
	}
	if val != "PONG" {
		return fmt.Errorf("Redis respond error %s %s", host, port)
	}
	return err
}
