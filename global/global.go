package global

import (
	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"os"
)

var (
	ws_host, ws_path, redis_host, redis_port string
	redis_conn                               redis.Conn
	logger                                   *log.Logger
	signals                                  chan os.Signal
)

func SetWsHost(wh string) {
	ws_host = wh
}

func SetWsPath(wp string) {
	ws_path = wp
}

func SetRedisHost(rh string) {
	redis_host = rh
}

func SetRedisPort(rp string) {
	redis_port = rp
}

func SetRedisConn(c redis.Conn) {
	redis_conn = c
}

func SetLogger(l *log.Logger) {
	logger = l
}

func SetSignals(s chan os.Signal) {
	signals = s
}

func Logger() *log.Logger {
	return logger
}

func RedisConn() redis.Conn {
	return redis_conn
}

func Signals() chan os.Signal {
	return signals
}

func WsHost() string {
	return ws_host
}

func WsPath() string {
	return ws_path
}

func RedisHost() string {
	return redis_host
}

func RedisPort() string {
	return redis_port
}
