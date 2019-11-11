package redis

import (
	"fmt"
	"log"
	netUrl "net/url"

	"github.com/tron-us/go-common/constant"

	redis "github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

type TGRDDB struct {
	rd *redis.Client
}

var (
	host = ""
	port = ""
	db   = 0
)

func CreateTGRDDB(url string) *TGRDDB {
	opts, err := netUrl.Parse(url)
	if err != nil {
		log.Panic(constant.DBURLParseError, zap.String("URL:", url), zap.Error(err))
	}
	password, _ := opts.User.Password()
	host = opts.Host
	port = opts.Port()
	var client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
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
