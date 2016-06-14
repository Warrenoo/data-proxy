package statistics

import (
	. "github.com/warrenoo/data-proxy/global"
	"github.com/warrenoo/data-proxy/models"
	"strconv"
	"syscall"
	"time"
)

var ServerKey = "realtime:server:" + strconv.Itoa(syscall.Getpid())

func InitStatistics() {
	Statistics["realtime:markets:updated_at"] = make(map[string]string)

	Statistics[ServerKey] = make(map[string]string)
	Statistics[ServerKey]["test"] = "123"

}

func UpdateStockUpdatedAt(stock *models.RealTimeStock) {
	if stock.Market != "" {
		Statistics["realtime:markets:updated_at"][stock.Market] = strconv.FormatInt(time.Now().Unix(), 10)
	}
}
