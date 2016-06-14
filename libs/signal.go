package libs

import (
	"fmt"
	. "github.com/warrenoo/data-proxy/global"
	"os/signal"
	"syscall"
	"time"
)

func RegisterSignal() {
	signal.Notify(Signals)

	fmt.Printf("Init Signal Ok\n")
}

func ListenSignal() {
	for {
		select {
		// 处理信号量
		case s := <-Signals:
			if s == syscall.SIGINT {
				Logger.Info("Get Signal: ", s)
				CloseFlag <- true

				// 阻塞2秒钟用来让其他线程结束
				time.Sleep(2 * time.Second)

				CloseAll()

				fmt.Print("Server Over!!\n")
				return
			}
		}
	}
}
