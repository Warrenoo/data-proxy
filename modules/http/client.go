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
	for _, m := range Markets {
		//InitBaseStock(m)
		InitRealTime(m)
	}
}

func InitDataByMarkets(markets []string) {
	for _, m := range markets {
		//InitBaseStock(m)
		InitRealTime(m)
	}
}

func InitRealTime(market string) {
	Logger.Info("Init RealTime " + market + " Begin!!")
	fmt.Printf("Init RealTime %s...\n", market)
	response_body := httpPost("http://d.caishuo.com:6090/websocket", "text/plain; charset=UTF-8", "{'channel':'9', 'market':'"+market+"'}")

	s := []string{}
	parseJson([]byte(response_body), &s)

	for _, v := range s {
		analysis.RealTimeStockPersistence(v)
	}

	fmt.Printf("Init RealTime %s Ok\n", market)
	Logger.Info("Init RealTime " + market + " Over!!")
}

func InitBaseStock(market string) {
	Logger.Info("Init BaseStock " + market + " Begin!!")
	fmt.Printf("Init BaseStock %s...\n", market)
	response_body := httpPost("http://d.caishuo.com:6090/websocket", "text/plain; charset=UTF-8", "{'channel':'9', 'market':'"+market+"'}")

	s := []string{}
	parseJson([]byte(response_body), &s)

	for _, v := range s {
		analysis.BaseStockPersistence(v)
	}

	fmt.Printf("Init BaseStock %s Ok\n", market)
	Logger.Info("Init BaseStock " + market + " Over!!")
}

func parseJson(data []byte, s *[]string) {
	json.Unmarshal(data, s)
}

func httpPost(host string, request_type string, message string) string {
	resp, err := http.Post(host, request_type, strings.NewReader(message))

	if err != nil {
		Logger.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Logger.Error(err)
	}

	return string(body)
}
