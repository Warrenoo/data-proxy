package libs

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"os"
)

func InitLog() {
	log := logrus.New()

	if f, err := os.OpenFile("go-data-client.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		fmt.Printf("opening file error: %v\n", err)
		// log to 标准输出
		log.Out = os.Stderr
	} else {
		// log to file
		log.Out = f
	}

	log.Level = logrus.InfoLevel

	// 输出logstash格式
	// 通过hook可以直接发送到logstash服务
	log.Formatter = &logstash.LogstashFormatter{Type: "go-data-client"}

	SetLogger(log)

	fmt.Printf("Init Log Ok\n")
}
