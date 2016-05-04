package main

import (
	"fmt"
	"gitlab.caishuo.com/ruby/go-data-client/libs"
	"gitlab.caishuo.com/ruby/go-data-client/modules"
)

func initConfig() {
	// 命令行参数解析
	libs.GetEnv()

	// 初始化redis连接
	libs.InitRedis()

	// 初始化log
	libs.InitLog()

	// 注册接收信号channel
	libs.RegisterSignal()
}

func main() {

	fmt.Printf("Start Worker......\n")

	initConfig()

	modules.InitData()

	// 开始websockst数据监听
	//modules.BeginWebSocket()

}
