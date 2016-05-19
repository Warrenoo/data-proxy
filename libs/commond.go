package libs

import (
	"flag"
)

func ParseCommond() (string, string, string, string, string, string, string, int, string, string) {
	ws_host := flag.String("dc_ws_host", "d.caishuo.com", "websocket host")
	ws_path := flag.String("dc_ws_path", "/websocket", "websocket path")
	ws_token := flag.String("dc_ws_token", "pro_caishuo_123!@#", "websocket token")
	ws_random := flag.String("dc_ws_random", "pro_caishuo", "websocket random")
	redis_host := flag.String("dc_redis_host", "127.0.0.1", "redis host")
	redis_port := flag.String("dc_redis_port", "6379", "redis port")
	redis_password := flag.String("dc_redis_password", "", "redis password")
	redis_database := flag.Int("dc_redis_database", 0, "redis database")
	http_server_port := flag.String("dc_http_server_port", "8000", "http server port")
	log_level := flag.String("dc_log_level", "WARN", "log level")

	flag.Parse()

	return *ws_host, *ws_path, *ws_token, *ws_random, *redis_host, *redis_port, *redis_password, *redis_database, *http_server_port, *log_level
}
