package analysis

import (
	"errors"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"gitlab.caishuo.com/ruby/go-data-client/models"
	"strings"
)

func analysisRealTimeStock(data string) ([]string, error) {
	str := strings.Replace(data, "|", ",", 1)
	str_arr := strings.Split(str, ",")

	if len(str_arr) != 19 {
		return nil, errors.New("data size must be 19\n")
	}

	return str_arr, nil
}

func makeRealTimeStock(data string) *models.Stock {
	result, err := analysisRealTimeStock(data)

	if err != nil {
		Logger().Error(err)
		return nil
	}

	return &models.Stock{result[0], result[1], result[2], result[3], result[4], result[5], result[6], result[7], result[8], result[9], result[10], result[11], result[12], result[13], result[14], result[15], result[16], result[17], result[18]}
}

func RealTimeStockPersistence(data string) *models.Stock {
	stock := makeRealTimeStock(data)

	if stock == nil {
		Logger().Error("stock make fail..., from data: ", data)
		return nil
	}

	Objects() <- stock

	return stock
}
