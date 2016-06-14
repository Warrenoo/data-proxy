package websocket

import (
	"github.com/warrenoo/data-proxy/analysis"
	. "github.com/warrenoo/data-proxy/global"
)

func baseStockHook(data string) interface{} {
	return analysis.BaseStockPersistence(data)
}

func StartBaseStock() {

	// 消息
	messages := make([]string, len(Markets))

	for i, m := range Markets {
		messages[i] = "{'isSubscribe':true,'market':'" + m + "','stock_code':'','channel':'9'}"
	}

	StartWebSocket(WsHost, WsPath, messages, baseStockHook)
}
