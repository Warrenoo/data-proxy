package libs

import (
	"fmt"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"os"
	"strconv"
)

func GetEnv() {
	ws_host, ws_path, ws_token, ws_random, redis_host, redis_port, redis_password, redis_database, http_server_port, log_level := os.Getenv("DC_WS_HOST"), os.Getenv("DC_WS_PATH"), os.Getenv("DC_WS_TOKEN"), os.Getenv("DC_WS_RANDOM"), os.Getenv("DC_REDIS_HOST"), os.Getenv("DC_REDIS_PORT"), os.Getenv("DC_REDIS_PASSWORD"), os.Getenv("DC_REDIS_DATABASE"), os.Getenv("DC_HTTP_SERVER_PORT"), os.Getenv("DC_LOG_LEVEL")
	ws_host_tmp, ws_path_tmp, ws_token_tmp, ws_random_tmp, redis_host_tmp, redis_port_tmp, redis_password_tmp, redis_database_tmp, http_server_port_tmp, log_level_tmp := ParseCommond()

	WsHost = andEqual(ws_host, ws_host_tmp).(string)
	WsPath = andEqual(ws_path, ws_path_tmp).(string)
	WsToken = andEqual(ws_token, ws_token_tmp).(string)
	WsRandom = andEqual(ws_random, ws_random_tmp).(string)
	RedisHost = andEqual(redis_host, redis_host_tmp).(string)
	RedisPort = andEqual(redis_port, redis_port_tmp).(string)
	RedisPassword = andEqual(redis_password, redis_password_tmp).(string)
	RedisDatabase = andEqual(redis_database, redis_database_tmp).(int)
	HttpServerPort = andEqual(http_server_port, http_server_port_tmp).(string)
	LogLevel = andEqual(log_level, log_level_tmp).(string)

	fmt.Printf("Init Commands Ok\n")
	fmt.Printf("Commands: ws_host=%s, ws_path=%s, redis_host=%s, redis_port=%s, redis_database=%d, http_server_port=%s, log_level=%s\n", WsHost, WsPath, RedisHost, RedisPort, RedisDatabase, HttpServerPort, LogLevel)
}

func andEqual(s1 string, s2 interface{}) interface{} {
	if s1 == "" {
		return s2
	} else {
		switch s2.(type) {
		case string:
			return s1
		case int:
			s, err := strconv.Atoi(s1)
			if err != nil {
				return s2
			} else {
				return s
			}
		}
	}
	return nil
}
