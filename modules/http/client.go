package http

import (
	"encoding/json"
	"fmt"
	"gitlab.caishuo.com/ruby/go-data-client/analysis"
	"io/ioutil"
	"net/http"
	"strings"
)

func InitData() {
	response_body := httpPost("http://d.caishuo.com:6090/websocket", "text/plain; charset=UTF-8", "{'channel':'9'}")

	s := []string{}
	json.Unmarshal([]byte(response_body), &s)

	for _, v := range s {
		analysis.StockPersistence(v)
		fmt.Printf("INIT DATA: %s\n", v)
	}
}

func httpPost(host string, request_type string, message string) string {
	resp, err := http.Post(host, request_type, strings.NewReader(message))

	if err != nil {
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
	}

	return string(body)
}
