package libs

import (
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	"os"
)

func InitLog() *logrus.Logger {
	log := logrus.New()

	log.Out = os.Stderr
	log.Level = logrus.InfoLevel

	// 输出logstash格式
	// 通过hook可以直接发送到logstash服务
	log.Formatter = &logstash.LogstashFormatter{Type: "go-data-client"}

	return log
}
