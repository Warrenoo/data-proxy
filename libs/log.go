package libs

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	. "github.com/warrenoo/data-proxy/global"
	"os"
	"strings"
)

func InitLog() {

	level_map := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	log := logrus.New()

	if f, err := os.OpenFile("go-data-client.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		fmt.Printf("opening file error: %v\n", err)
		// log to 标准输出
		log.Out = os.Stderr
	} else {
		// log to file
		log.Out = f
		LogFile = f
	}

	log.Level = level_map[strings.ToLower(LogLevel)]

	// 输出logstash格式
	// 通过hook可以直接发送到logstash服务
	log.Formatter = &logstash.LogstashFormatter{Type: "go-data-client"}

	Logger = log

	fmt.Printf("Init Log Ok\n")
}
