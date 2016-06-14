package websocket

import (
	"github.com/warrenoo/data-prox/analysis"
	. "github.com/warrenoo/data-prox/global"
)

func realTimeHook(data string) interface{} {
	return analysis.RealTimeStockPersistence(data)
}

func StartRealtime() {

	// 消息
	messages := make([]string, len(Markets))

	for i, m := range Markets {
		messages[i] = "{'isSubscribe':true,'market':'" + m + "','stock_code':'','channel':'9','token':'" + WsToken + "','random':'" + WsRandom + "'}"
	}

	StartWebSocket(WsHost, WsPath, messages, realTimeHook)
}
