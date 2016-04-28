package main

import (
	"fmt"
	"github.com/warrenoo/getter"
	"gitlab.caishuo.com/ruby/go-data-client/analysis"
	"gitlab.caishuo.com/ruby/go-data-client/libs"
	"syscall"
	"time"
)

func main() {

	fmt.Printf("Start Worker......\n")

	// 命令行参数解析
	ws_host, ws_path, redis_host, redis_port := libs.GetEnv()
	fmt.Printf("command: ws_host=%s, ws_path=%s, redis_host=%s, redis_port=%s\n", ws_host, ws_path, redis_host, redis_port)

	// 初始化redis连接
	conn := libs.InitRedis(redis_host, redis_port)

	// 初始化log
	log := libs.InitLog()

	// 注册接收信号channel
	signals := libs.RegisterSignal()

	var url string = ws_host
	var path string = ws_path

	client := getter.New(url, path)

	// 连接websocket
	client.OnOpen(func() {
		log.Info("Init: ", url+path)
	})

	// 消息
	message := []byte("{'isSubscribe':true,'market':'hk','stock_code':'','channel':'9'}")
	log.Info("Listen: ", string(message))

	counter := 0
	// 发送一个消息，然后监听返回
	client.OnListen(message, func() {
		for {
			select {

			// 处理返回结果
			case data := <-client.Ch():

				// TODO 解析data
				stock := analysis.Make1(data.Data())
				if stock == nil {
					continue
				}

				counter++
				log.Info("Receive(", counter, "): ", stock, "cost_time: ", time.Now().Sub(data.St()))
				conn.Do("SET", "websock_test:"+time.Now().String(), stock.SaveFormat())

			// 处理信号量
			case s := <-signals:

				if s == syscall.SIGKILL || s == syscall.SIGINT {
					log.Info("get signal: ", s)
					client.Done() <- true
					return
				}
			}
		}
	})

	defer client.OnClose(func() {
		log.Info("Close Server!!")
	})
}
