package analysis

import (
	"errors"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"gitlab.caishuo.com/ruby/go-data-client/models"
	"strings"
)

func analysisBaseStock(data string) ([]string, error) {
	str := strings.Replace(data, "|", ",", 1)
	str_arr := strings.Split(str, ",")

	if len(str_arr) != 19 {
		return nil, errors.New("data size must be 19\n")
	}

	return str_arr, nil
}

func makeBaseStock(data string) *models.Stock {
	result, err := analysisBaseStock(data)

	if err != nil {
		Logger().Error(err)
		return nil
	}

	return &models.Stock{result[0], result[1], result[2], result[3], result[4], result[5], result[6], result[7], result[8], result[9], result[10], result[11], result[12], result[13], result[14], result[15], result[16], result[17], result[18]}
}

func BaseStockPersistence(data string) *models.Stock {
	stock := makeBaseStock(data)

	if stock == nil {
		Logger().Error("stock make fail..., from data: ", data)
		return nil
	}

	Objects() <- stock

	return stock
}
