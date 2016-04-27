package libs

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func InitRedis() redis.Conn {
	conn, _ := redis.DialTimeout("tcp", "192.168.99.100:6379", 0, 1*time.Second, 1*time.Second)
	return conn
}
