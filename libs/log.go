package libs

import (
	//"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	"os"
)

func InitLog() *logrus.Logger {
	log := logrus.New()

	// log to a file
	//f, err := os.OpenFile("testlogrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	//if err != nil {
	//fmt.Printf("error opening file: %v", err)
	//}
	//log.Out = f

	// log to 标准输出
	log.Out = os.Stderr

	log.Level = logrus.InfoLevel

	// 输出logstash格式
	// 通过hook可以直接发送到logstash服务
	log.Formatter = &logstash.LogstashFormatter{Type: "go-data-client"}

	return log
}
