package http

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"net/http"
	"os"
	"syscall"
)

func info(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO Data CLIENT running......(pid: %d)\n", syscall.Getpid())
	fmt.Fprintf(w, "Commands: ws_host=%s, ws_path=%s, redis_host=%s, redis_port=%s, redis_database=%d, http_server_port=%s, log_level=%s\n", WsHost, WsPath, RedisHost, RedisPort, RedisDatabase, HttpServerPort, LogLevel)

	fmt.Fprintf(w, "-------------------统计数据------------------\n")
	for key, value := range Statistics {
		fmt.Fprintf(w, "- %s:\n", key)
		for k, v := range value {
			fmt.Fprintf(w, "-- %s: %s\n", k, v)
		}
	}
}

func resend(c web.C, w http.ResponseWriter, r *http.Request) {

	market := c.URLParams["market"]
	go InitDataByMarkets([]string{market})

	fmt.Fprintf(w, "已经开始同步%s市场股票数据!", market)
}

func stock_id(c web.C, w http.ResponseWriter, r *http.Request) {

	symbol := c.URLParams["symbol"]

	fmt.Fprintf(w, "Symbol: %s ----> Id: %s", symbol, IdsMap[symbol])
}

func StartWebServer() {
	os.Setenv("PORT", HttpServerPort)
	goji.Get("/info", info)
	goji.Get("/resend/:market", resend)
	goji.Get("/stock_id/:symbol", stock_id)

	goji.Serve()
}
