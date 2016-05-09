package websocket

import (
	"gitlab.caishuo.com/ruby/go-data-client/analysis"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
)

func baseStockHook(data string) interface{} {
	return analysis.BaseStockPersistence(data)
}

func StartBaseStock() {

	// 消息
	var messages = []string{
		"{'isSubscribe':true,'market':'hk','stock_code':'','channel':'9'}",
		"{'isSubscribe':true,'market':'sh','stock_code':'','channel':'9'}",
		"{'isSubscribe':true,'market':'sz','stock_code':'','channel':'9'}",
		"{'isSubscribe':true,'market':'us','stock_code':'','channel':'9'}",
	}

	StartWebSocket(WsHost(), WsPath(), messages, baseStockHook)
}
