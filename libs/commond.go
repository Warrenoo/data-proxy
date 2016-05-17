package libs

import (
	"flag"
)

func ParseCommond() (string, string, string, string, string, string) {
	ws_host := flag.String("dc_ws_host", "d.caishuo.com", "websocket host")
	ws_path := flag.String("dc_ws_path", "/websocket", "websocket path")
	redis_host := flag.String("dc_redis_host", "127.0.0.1", "redis host")
	redis_port := flag.String("dc_redis_port", "6379", "redis port")
	redis_password := flag.String("dc_redis_password", "", "redis password")
	http_server_port := flag.String("dc_http_server_port", "8000", "http server port")

	flag.Parse()

	return *ws_host, *ws_path, *redis_host, *redis_port, *redis_password, *http_server_port
}
