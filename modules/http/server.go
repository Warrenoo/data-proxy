package http

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"net/http"
	"os"
)

func info(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO Data CLIENT running......\n")
	fmt.Fprintf(w, "commands: ws_host=%s, ws_path=%s, redis_host=%s, redis_port=%s, http_server_port=%s\n", WsHost(), WsPath(), RedisHost(), RedisPort(), HttpServerPort())
}

func resend(c web.C, w http.ResponseWriter, r *http.Request) {
	go InitData()
	fmt.Fprintf(w, "已经开始同步全部股票数据!")
}

func StartWebServer() {
	os.Setenv("PORT", HttpServerPort())
	goji.Get("/info", info)
	goji.Post("/resend", resend)

	goji.Serve()
}
