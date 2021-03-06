package analysis

import (
	"errors"
	. "github.com/warrenoo/data-proxy/global"
	"github.com/warrenoo/data-proxy/models"
	"strings"
)

func analysisBaseStock(data string) ([]string, error) {
	str_arr := strings.Split(data, "~")

	if len(str_arr) < 7 {
		return nil, errors.New("data size must be >= 7\n")
	}

	return str_arr, nil
}

func makeBaseStock(data string) *models.BaseStock {
	result, err := analysisBaseStock(data)

	if err != nil {
		Logger.Error(err)
		return nil
	}

	return &models.BaseStock{Symbol: result[0], Id: result[1], Exchange: result[2], Naps: result[3], Eps: result[4], MarketCapitalization: result[5], PeRatio: result[6]}
}

func BaseStockPersistence(data string) *models.BaseStock {
	stock := makeBaseStock(data)

	mapSymbolAndId(stock)

	if stock == nil {
		Logger.Error("stock make fail..., from data: ", data)
		return nil
	}

	Objects <- stock

	return stock
}

func mapSymbolAndId(stock *models.BaseStock) {
	IdsMap[stock.Symbol] = stock.Id
}
