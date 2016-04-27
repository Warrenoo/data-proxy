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

	var url string = "d.caishuo.com:6090"
	var path string = "/websocket"

	client := getter.New(url, path)

	// 连接websocket
	client.OnOpen(func() {
		fmt.Printf("Init: %s\n", url)
	})

	// 消息1
	message := []byte("{'isSubscribe':true,'market':'hk','stock_code':'','channel':'9'}")
	fmt.Printf("Listen: %s\n", message)

	counter := 0

	// 发送一个消息，然后监听返回
	client.OnListen(message, func() {
		for {
			select {

			// 处理返回结果
			case data := <-client.Ch():

				// TODO 解析data

				counter++
				fmt.Printf("pid: %d\n", os.Getpid())
				fmt.Printf("Receive(%d): %s\n", counter, data)
				conn.Do("SET", "websock_test:"+time.Now().String(), data)

			// 处理信号量
			case s := <-signals:

				if s == syscall.SIGKILL || s == syscall.SIGINT {
					fmt.Println("get signal:", s)
					client.Done() <- true
					return
				}
			}
		}
	})

	defer client.OnClose(func() {
		fmt.Printf("Close Server!!")
	})
}
