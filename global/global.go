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
	log_file         *os.File
	signals          chan os.Signal
	close_flag       chan bool
	ids_map          map[string]string
	ws_host          string
	ws_path          string
	redis_host       string
	redis_port       string
	http_server_port string
)

func CloseAll() {
	close(signals)
	close(close_flag)
	close(objects)
	log_file.Close()
}

func SetCloseFlag(cf chan bool) {
	close_flag = cf
}

func SetIdsMap(im map[string]string) {
	ids_map = im
}

func SetObjects(o chan interface{}) {
	objects = o
}

func SetLogFile(f *os.File) {
	log_file = f
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

func LogFile() *os.File {
	return log_file
}

func IdsMap() map[string]string {
	return ids_map
}

func CloseFlag() chan bool {
	return close_flag
}
