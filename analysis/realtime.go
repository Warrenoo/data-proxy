package analysis

import (
	"errors"
	. "github.com/warrenoo/data-proxy/global"
	"github.com/warrenoo/data-proxy/models"
	"github.com/warrenoo/data-proxy/statistics"
	"strings"
)

func analysisRealTimeStock(data string) ([]string, error) {

	str := strings.Replace(data, "|", ",", 1)
	str_arr := strings.Split(str, ",")

	if len(str_arr) < 22 {
		return nil, errors.New("data size must >= 22\n")
	}

	return str_arr, nil
}

func makeRealTimeStock(data string) *models.RealTimeStock {
	result, err := analysisRealTimeStock(data)

	if err != nil {
		Logger.Error(err)
		return nil
	}

	stock := models.RealTimeStock{
		Channel:         result[0],
		Market:          result[1],
		TradeTime:       result[2],
		Symbol:          result[3],
		RealTimePrice:   result[4],
		ChangePrice:     result[5],
		ChangePriceRate: result[6],
		TodayHigh:       result[7],
		TodayLow:        result[8],
		Volume:          result[9],
		Buy1Price:       result[10],
		Buy1Num:         result[11],
		Sell1Price:      result[12],
		Sell1Num:        result[13],
		Week52High:      result[14],
		Week52Low:       result[15],
		Amount:          result[16],
		TodayOpen:       result[17],
		LastClose:       result[18],
		Amplitude:       result[19],
		TopPrice:        result[20],
		BottomPrice:     result[21],
	}

	if stock.Market == "sz" || stock.Market == "sh" {
		stock.Symbol = stock.Symbol + "." + strings.ToUpper(stock.Market)
	}
	id := IdsMap[stock.Symbol]

	if id != "" {
		stock.Id = id
	} else {
		return nil
		//stock.Id = stock.Symbol
	}

	return &stock
}

func RealTimeStockPersistence(data string) *models.RealTimeStock {
	stock := makeRealTimeStock(data)

	if stock == nil {
		Logger.Error("stock make fail..., from data: ", data)
		return nil
	}

	Objects <- stock
	statistics.UpdateStockUpdatedAt(stock)

	return stock
}
