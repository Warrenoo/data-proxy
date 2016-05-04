package libs

import (
	"fmt"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"os"
)

func GetEnv() {
	ws_host, ws_path, redis_host, redis_port := os.Getenv("DC_WS_HOST"), os.Getenv("DC_WS_PATH"), os.Getenv("DC_REDIS_HOST"), os.Getenv("DC_REDIS_PORT")
	ws_host_tmp, ws_path_tmp, redis_host_tmp, redis_port_tmp := ParseCommond()

	SetWsHost(andEqual(ws_host, ws_host_tmp))
	SetWsPath(andEqual(ws_path, ws_path_tmp))
	SetRedisHost(andEqual(redis_host, redis_host_tmp))
	SetRedisPort(andEqual(redis_port, redis_port_tmp))

	fmt.Printf("commands: ws_host=%s, ws_path=%s, redis_host=%s, redis_port=%s\n", WsHost(), WsPath(), RedisHost(), RedisPort())
}

func andEqual(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	} else {
		return s1
	}
}
