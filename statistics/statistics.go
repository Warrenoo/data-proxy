package statistics

import (
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"gitlab.caishuo.com/ruby/go-data-client/models"
	"strconv"
	"time"
)

func InitStatistics() {
	Statistics["realtime:markets:updated_at"] = make(map[string]string)
}

func UpdateStockUpdatedAt(stock *models.Stock) {
	Statistics["realtime:markets:updated_at"][stock.Market] = strconv.FormatInt(time.Now().Unix(), 10)
}
