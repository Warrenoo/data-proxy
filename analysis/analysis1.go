package analysis

import (
	"errors"
	"fmt"
	"gitlab.caishuo.com/ruby/go-data-client/models"
	"strings"
)

func AnalysisChannel1(data string) ([]string, error) {
	str := strings.Replace(data, "|", ",", 1)
	str_arr := strings.Split(str, ",")

	if len(str_arr) != 18 {
		return nil, errors.New("data size must be 18")
	}

	return str_arr, nil
}

func Make1(data string) *models.Stock {
	result, err := AnalysisChannel1(data)

	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	return &models.Stock{result[0], result[1], result[2], result[3], result[4], result[5], result[6], result[7], result[8], result[9], result[10], result[11], result[12], result[13], result[14], result[15], result[16], result[17]}
}