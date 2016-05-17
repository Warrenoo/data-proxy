package libs

import (
	"fmt"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"os"
)

func GetEnv() {
	ws_host, ws_path, redis_host, redis_port, redis_password, http_server_port := os.Getenv("DC_WS_HOST"), os.Getenv("DC_WS_PATH"), os.Getenv("DC_REDIS_HOST"), os.Getenv("DC_REDIS_PORT"), os.Getenv("DC_REDIS_PASSWORD"), os.Getenv("DC_HTTP_SERVER_PORT")
	ws_host_tmp, ws_path_tmp, redis_host_tmp, redis_port_tmp, redis_password_tmp, http_server_port_tmp := ParseCommond()

	WsHost = andEqual(ws_host, ws_host_tmp)
	WsPath = andEqual(ws_path, ws_path_tmp)
	RedisHost = andEqual(redis_host, redis_host_tmp)
	RedisPort = andEqual(redis_port, redis_port_tmp)
	RedisPassword = andEqual(redis_password, redis_password_tmp)
	HttpServerPort = andEqual(http_server_port, http_server_port_tmp)

	fmt.Printf("Init Commands Ok\n")
	fmt.Printf("Commands: ws_host=%s, ws_path=%s, redis_host=%s, redis_port=%s, http_server_port=%s\n", WsHost, WsPath, RedisHost, RedisPort, HttpServerPort)
}

func andEqual(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	} else {
		return s1
	}
}
