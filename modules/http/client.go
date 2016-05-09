package http

import (
	"encoding/json"
	"fmt"
	"gitlab.caishuo.com/ruby/go-data-client/analysis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"io/ioutil"
	"net/http"
	"strings"
)

func InitData() {
	//InitBaseStock()
	InitRealTime()
}

func InitRealTime() {
	Logger().Info("Init RealTime Begin!!")
	fmt.Printf("Init RealTime...\n")
	response_body := httpPost("http://d.caishuo.com:6090/websocket", "text/plain; charset=UTF-8", "{'channel':'9'}")

	s := []string{}
	parseJson([]byte(response_body), &s)

	for _, v := range s {
		analysis.RealTimeStockPersistence(v)
	}

	fmt.Printf("Init RealTime Ok\n")
	Logger().Info("Init RealTime Over!!")
}

func InitBaseStock() {
	Logger().Info("Init BaseStock Begin!!")
	fmt.Printf("Init BaseStock...\n")
	response_body := httpPost("http://d.caishuo.com:6090/websocket", "text/plain; charset=UTF-8", "{'channel':'9'}")

	s := []string{}
	parseJson([]byte(response_body), &s)

	for _, v := range s {
		analysis.BaseStockPersistence(v)
	}

	fmt.Printf("Init BaseStock Ok\n")
	Logger().Info("Init BaseStock Over!!")
}

func parseJson(data []byte, s *[]string) {
	json.Unmarshal(data, s)
}

func httpPost(host string, request_type string, message string) string {
	resp, err := http.Post(host, request_type, strings.NewReader(message))

	if err != nil {
		Logger().Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Logger().Error(err)
	}

	return string(body)
}
