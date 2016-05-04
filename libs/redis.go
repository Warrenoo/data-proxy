package libs

import (
	"github.com/garyburd/redigo/redis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"time"
)

func InitRedis() {
	conn, err := redis.DialTimeout("tcp", RedisHost()+":"+RedisPort(), 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err.Error())
	}

	SetRedisConn(conn)
}
