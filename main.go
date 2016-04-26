package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/warrenoo/getter"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 初始化redis连接
	conn, _ := redis.DialTimeout("tcp", "192.168.99.100:6379", 0, 1*time.Second, 1*time.Second)

	// 信号处理
	signals := make(chan os.Signal)
	signal.Notify(signals)

	var origin string = "http://douniwan.com"
	var url string = "ws://d.caishuo.com:6090/websocket"

	client := getter.New(origin, url, "")

	// 连接websocket
	client.OnConnect(func() {
		fmt.Printf("Init: %s\n", url)
	})

	// 消息
	message := getter.Data("{'isSubscribe':true,'market':'hk','stock_code':'00700','channel':'1,2,3'}")
	fmt.Printf("Listen: %s\n", message)

	// 发送一个消息，然后监听返回
	client.OnListen(message, func() {
		for {
			select {

			// 处理返回结果
			case data := <-client.Ch():

				// TODO 解析data

				fmt.Printf("Receive: %s\n", data)
				conn.Do("SET", "websock_test:"+time.Now().String(), data)

			// 处理信号量
			case s := <-signals:

				fmt.Println("get signal:", s)

				if s == syscall.SIGKILL || s == syscall.SIGINT {
					client.Done() <- true
					return
				}
			}
		}
	})

	client.OnClose(func() {
		fmt.Printf("Close Server!!")
	})
}
