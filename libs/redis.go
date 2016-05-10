package libs

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"gitlab.caishuo.com/ruby/go-data-client/models"
	"time"
)

func InitRedis() {
	conn, err := redis.DialTimeout("tcp", RedisHost()+":"+RedisPort(), 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err)
	}

	objects := make(chan interface{}, 1024)

	SetObjects(objects)
	SetRedisConn(conn)

	fmt.Printf("Init Redis Ok\n")

	go redisWrite()
	fmt.Printf("RedisWrite Running...\n")
}

func redisWrite() {
	for {
		select {
		case o := <-Objects():

			switch o.(type) {
			case *models.Stock:

				stock := o.(*models.Stock)
				if _, err := RedisConn().Do("HMSET", redis.Args{}.Add("realtime:"+stock.Id).AddFlat(stock)...); err != nil {
					Logger().Error("Redis Error: ", err)
				}

			}

		case <-CloseFlag():
			CloseFlag() <- true

			RedisConn().Close()
			fmt.Printf("Redis Conn Closed!\n")
			return
		}
	}
}
