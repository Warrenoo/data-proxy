package main

import (
	"fmt"
	. "github.com/warrenoo/data-proxy/libs"
	. "github.com/warrenoo/data-proxy/modules/http"
	. "github.com/warrenoo/data-proxy/modules/websocket"
	. "github.com/warrenoo/data-proxy/statistics"
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

	// 初始化统计信息
	InitStatistics()
}

func main() {

	fmt.Printf("Start Worker......\n")

	initConfig()

	// 开启web服务
	go StartWebServer()

	// 初始化数据
	InitData()

	// 开始websocket监听
	//go StartBaseStock()
	go StartRealtime()

	// 阻塞监听信号
	ListenSignal()
}
