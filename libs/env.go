package libs

import (
	"os"
)

func GetEnv() (string, string, string, string) {
	ws_host, ws_path, redis_host, redis_port := os.Getenv("DC_WS_HOST"), os.Getenv("DC_WS_PATH"), os.Getenv("DC_REDIS_HOST"), os.Getenv("DC_REDIS_PORT")
	ws_host_tmp, ws_path_tmp, redis_host_tmp, redis_port_tmp := ParseCommond()

	return andEqual(ws_host, ws_host_tmp), andEqual(ws_path, ws_path_tmp), andEqual(redis_host, redis_host_tmp), andEqual(redis_port, redis_port_tmp)
}

func andEqual(s1 string, s2 string) string {
	if s1 == "" {
		return s2
	} else {
		return s1
	}
}
