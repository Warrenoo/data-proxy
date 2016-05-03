package libs

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func InitRedis(redis_host string, redis_port string) redis.Conn {
	conn, err := redis.DialTimeout("tcp", redis_host+":"+redis_port, 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
