package global

import (
	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"os"
)

var (
	RedisConn      redis.Conn
	Logger         *log.Logger
	LogFile        *os.File
	WsHost         string
	WsPath         string
	WsToken        string
	WsRandom       string
	RedisHost      string
	RedisPort      string
	RedisPassword  string
	RedisDatabase  int
	HttpServerPort string
	LogLevel       string
)

var (
	Statistics = make(map[string]map[string]string)
	IdsMap     = make(map[string]string)
	Objects    = make(chan interface{}, 1024)
	Signals    = make(chan os.Signal)
	CloseFlag  = make(chan bool, 5)
	Markets    = []string{"hk", "cn", "us"}
)

func CloseAll() {
	close(Signals)
	close(CloseFlag)
	close(Objects)
	LogFile.Close()
}
