package common

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisConf struct {
	Addr        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

func InitRedis(conf RedisConf) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Addr)
		},
	}

	conn := pool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	return
}
