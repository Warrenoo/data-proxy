package modules

import (
	"fmt"
	"github.com/warrenoo/getter"
	"gitlab.caishuo.com/ruby/go-data-client/analysis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"syscall"
	"time"
)

func StartWebSocket() {
	fmt.Printf("Init WebSocket...\n")
	client := getter.New(WsHost(), WsPath())

	// 连接websocket
	client.OnOpen(func() {
		Logger().Info("Connect WebSocket: ", WsHost()+WsPath())
	})

	// 消息
	var messages = []string{
		"{'isSubscribe':true,'market':'hk','stock_code':'','channel':'9'}",
		"{'isSubscribe':true,'market':'sh','stock_code':'','channel':'9'}",
		"{'isSubscribe':true,'market':'sz','stock_code':'','channel':'9'}",
		"{'isSubscribe':true,'market':'us','stock_code':'','channel':'9'}",
	}

	// 发送一个消息，然后监听返回
	client.OnListen(&messages, func() {
		for {
			select {

			// 处理返回结果
			case data := <-client.Ch():

				if stock := analysis.StockPersistence(data.Data()); stock == nil {
					continue
				} else {
					Logger().Info("Receive: ", stock, ", cost_time: ", time.Now().Sub(data.St()))
				}

			// 处理信号量
			case s := <-Signals():

				if s == syscall.SIGINT {
					Logger().Info("Get Signal: ", s)
					return
				}
			}
		}
	})

	defer client.OnClose(func() {
		fmt.Printf("Close Redis...\n")
		RedisConn().Close()
		fmt.Printf("Close LogFile...\n")
		LogFile().Close()
		fmt.Printf("Close Server Over!!\n")
	})

}
