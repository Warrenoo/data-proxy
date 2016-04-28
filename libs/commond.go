package libs

import (
	"flag"
	"fmt"
)

func ParseCommond() (string, string, string, string) {
	ws_host := flag.String("DC_WS_HOST", "d.caishuo.com", "websocket host")
	ws_path := flag.String("DC_WS_PATH", "/websocket", "websocket path")
	redis_host := flag.String("DC_REDIS_HOST", "127.0.0.1", "redis host")
	redis_port := flag.String("DC_REDIS_PORT", "6379", "redis port")

	flag.Parse()

	fmt.Printf("commonds: DC_WS_HOST=%s, DC_WS_PATH=%s, DC_REDIS_HOST=%s, DC_REDIS_PORT=%s\n", *ws_host, *ws_path, *redis_host, *redis_port)

	return *ws_host, *ws_path, *redis_host, *redis_port
}
