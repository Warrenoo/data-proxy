package websocket

import (
	"github.com/warrenoo/getter"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"time"
)

func StartWebSocket(host string, path string, messages []string, listen_hook func(string) interface{}) {
	client := getter.New(host, path)

	// 连接websocket
	client.OnOpen(func(this *getter.Client) {
		Logger().Info("Connect WebSocket: ", host+path)
	})

	// 发送一个消息，然后监听返回
	client.OnListen(&messages, func(client *getter.Client) {
		for {
			select {
			// 处理返回结果
			case data := <-client.Ch():

				if obj := listen_hook(data.Data()); obj == nil {
					continue
				} else {
					Logger().Info("Receive: ", obj, ", cost_time: ", time.Now().Sub(data.St()))
				}

			case <-CloseFlag():
				CloseFlag() <- true
				client.Done() <- true

				time.Sleep(1 * time.Second)
				return
			}
		}

	})

	defer client.OnClose(func(this *getter.Client) {
		Logger().Info("WebSocket ", host+path, " Closed!")
	})
}
