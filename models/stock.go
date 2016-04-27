package models

import (
	"strings"
)

type Stock struct {
	Channel         string
	Market          string
	TradeTime       string
	Symbol          string
	RealTimePrice   string
	ChangePrice     string
	ChangePriceRate string
	TodayHigh       string
	TodayLow        string
	Volume          string
	Buy1Price       string
	Buy1Num         string
	Sell1Price      string
	Sell1Num        string
	Week52High      string
	Week52Low       string
	TodayOpen       string
	LastClose       string
}

func (this *Stock) SaveFormat() string {
	str := []string{this.Channel, this.Market, this.TradeTime, this.Symbol, this.RealTimePrice}
	return strings.Join(str, "--")
}
