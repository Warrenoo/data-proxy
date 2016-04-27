package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	"github.com/garyburd/redigo/redis"
	"github.com/warrenoo/getter"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init_log() *logrus.Logger {
	log := logrus.New()

	log.Out = os.Stderr
	log.Level = logrus.InfoLevel

	// 输出logstash格式
	// 通过hook可以直接发送到logstash服务
	log.Formatter = &logstash.LogstashFormatter{Type: "go-data-client"}

	return log
}

func init_redis() redis.Conn {
	conn, _ := redis.DialTimeout("tcp", "192.168.99.100:6379", 0, 1*time.Second, 1*time.Second)
	return conn
}

func main() {

	// 初始化redis连接
	conn := init_redis()

	// log 初始化
	log := init_log()

	// 信号处理
	signals := make(chan os.Signal)
	signal.Notify(signals)

	var url string = "d.caishuo.com:6090"
	var path string = "/websocket"

	client := getter.New(url, path)

	// 连接websocket
	client.OnOpen(func() {
		log.Info("Init: ", url+path)
	})

	// 消息1
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

				counter++
				log.Info("Receive(", counter, "): ", data)
				conn.Do("SET", "websock_test:"+time.Now().String(), data)

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
