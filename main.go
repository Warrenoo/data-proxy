package main

import (
	"fmt"
	. "gitlab.caishuo.com/ruby/go-data-client/libs"
	. "gitlab.caishuo.com/ruby/go-data-client/modules"
	. "gitlab.caishuo.com/ruby/go-data-client/modules/http"
)

func initConfig() {
	// 命令行参数解析
	GetEnv()

	// 初始化redis连接
	InitRedis()

	// 初始化log
	InitLog()

	// 注册接收信号channel
	RegisterSignal()
}

func main() {

	fmt.Printf("Start Worker......\n")

	initConfig()

	// 开启web服务
	go StartWebServer()

	// 初始化数据
	go InitData()

	// 开始websockst数据阻塞式监听
	StartWebSocket()

}
