package modules

import (
	"fmt"
	"github.com/warrenoo/getter"
	"gitlab.caishuo.com/ruby/go-data-client/analysis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"syscall"
	"time"
)

func BeginWebSocket() {
	client := getter.New(WsHost(), WsPath())

	// 连接websocket
	client.OnOpen(func() {
		Logger().Info("Connect WebSocket: ", WsHost()+WsPath())
	})

	// 消息
	var messages = []string{
		"{'isSubscribe':true,'market':'hk','stock_code':'','channel':'9'}",
		//"{'isSubscribe':true,'market':'sh','stock_code':'','channel':'9'}",
		//"{'isSubscribe':true,'market':'sz','stock_code':'','channel':'9'}",
		//"{'isSubscribe':true,'market':'us','stock_code':'','channel':'9'}",
	}

	// 发送一个消息，然后监听返回
	client.OnListen(&messages, func() {
		for {
			select {

			// 处理返回结果
			case data := <-client.Ch():
				stock := analysis.StockPersistence(data.Data())
				if stock == nil {
					continue
				}

				Logger().Info("Receive: ", &stock, ", cost_time: ", time.Now().Sub(data.St()))

			// 处理信号量
			case s := <-Signals():

				if s == syscall.SIGKILL || s == syscall.SIGINT {
					fmt.Printf("get signal: %s\n", s)
					client.Done() <- true
					return
				}
			}
		}
	})

	defer client.OnClose(func() {
		fmt.Printf("Close Server!!\n")
	})

}
