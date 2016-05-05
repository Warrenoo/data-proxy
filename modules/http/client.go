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
	Logger().Info("Init Data Begin!!")
	fmt.Printf("Init Data...\n")
	response_body := httpPost("http://d.caishuo.com:6090/websocket", "text/plain; charset=UTF-8", "{'channel':'9'}")

	s := []string{}
	json.Unmarshal([]byte(response_body), &s)

	for _, v := range s {
		analysis.StockPersistence(v)
	}

	fmt.Printf("Init Data Ok\n")
	Logger().Info("Init Data Over!!")
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
