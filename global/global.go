package global

import (
	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"os"
)

var (
	redis_conn       redis.Conn
	objects          chan interface{}
	logger           *log.Logger
	signals          chan os.Signal
	ws_host          string
	ws_path          string
	redis_host       string
	redis_port       string
	http_server_port string
)

func SetObjects(o chan interface{}) {
	objects = o
}

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

func SetHttpServerPort(hsp string) {
	http_server_port = hsp
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

func HttpServerPort() string {
	return http_server_port
}

func Objects() chan interface{} {
	return objects
}
