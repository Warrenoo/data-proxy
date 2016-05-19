package libs

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"gitlab.caishuo.com/ruby/go-data-client/models"
	"time"
)

func InitRedis() {
	conn, err := redis.Dial("tcp", RedisHost+":"+RedisPort, redis.DialConnectTimeout(0), redis.DialReadTimeout(1*time.Second), redis.DialWriteTimeout(1*time.Second), redis.DialPassword(RedisPassword), redis.DialDatabase(RedisDatabase))

	if err != nil {
		panic(err)
	}

	RedisConn = conn

	fmt.Printf("Init Redis Ok\n")

	go redisWrite()
	fmt.Printf("RedisWrite Running...\n")
}

func redisWrite() {

	ticker := time.NewTicker(60 * time.Second)

	for {
		select {
		case o := <-Objects:

			switch o.(type) {
			case *models.Stock:

				stock := o.(*models.Stock)
				hmset("realtime:"+stock.Id, stock)

			case map[string]map[string]string:
				for key, value := range o.(map[string]map[string]string) {
					hmset(key, value)
				}
			}

		// 每60秒写一次统计数据
		case <-ticker.C:
			Objects <- Statistics

		case <-CloseFlag:
			CloseFlag <- true

			RedisConn.Close()
			fmt.Printf("Redis Conn Closed!\n")
			return
		}
	}
}

func hmset(key string, obj interface{}) {
	if _, err := RedisConn.Do("HMSET", redis.Args{}.Add(key).AddFlat(obj)...); err != nil {
		Logger.Error("Redis Error: ", err)
	}
}
