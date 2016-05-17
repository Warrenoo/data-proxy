package statistics

import (
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"gitlab.caishuo.com/ruby/go-data-client/models"
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

func UpdateStockUpdatedAt(stock *models.Stock) {
	if stock.Market != "" {
		Statistics["realtime:markets:updated_at"][stock.Market] = strconv.FormatInt(time.Now().Unix(), 10)
	}
}
