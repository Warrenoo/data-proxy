package models

import ()

type Stock struct {
	Channel         string `redis:"channel"`
	Market          string `redis:"market"`
	TradeTime       string `redis:"trade_time"`
	Symbol          string `redis:"symbol"`
	RealTimePrice   string `redis:"last"`
	ChangePrice     string `redis:"change_from_previous_close"`
	ChangePriceRate string `redis:"percent_change_from_previous_close"`
	TodayHigh       string `redis:"high"`
	TodayLow        string `redis:"low"`
	Volume          string `redis:"volume"`
	Buy1Price       string `redis:"bid_prices"`
	Buy1Num         string `redis:"bid_sizes"`
	Sell1Price      string `redis:"offer_prices"`
	Sell1Num        string `redis:"offer_sizes"`
	Week52High      string `redis:"high52_weeks"`
	Week52Low       string `redis:"low52_weeks"`
	Amount          string `redis:"total_value_trade"`
	TodayOpen       string `redis:"open"`
	LastClose       string `redis:"previous_close"`
}

func (this *Stock) SaveFormat() *map[string]string {
	s := map[string]string{
		"high52_weeks":                       this.Week52High,
		"low52_weeks":                        this.Week52Low,
		"market_capitalization":              "",
		"pe_ratio":                           "",                   // 市盈率
		"exchange":                           "",                   // 交易所
		"symbol":                             this.Symbol,          // 股票代码
		"market":                             this.Market,          // 市场
		"trade_time":                         this.TradeTime,       // 交易时间
		"previous_close":                     this.LastClose,       // 昨收
		"open":                               this.TodayOpen,       // 今开
		"high":                               this.TodayHigh,       // 最高
		"low":                                this.TodayLow,        // 最低
		"last":                               this.RealTimePrice,   // 实时价格
		"volume":                             this.Volume,          // 成交量
		"total_value_trade":                  this.Amount,          // 成交额
		"bid_sizes":                          this.Buy1Num,         // 买一到买五数量逗号分隔
		"bid_prices":                         this.Buy1Price,       // 买一到买五价格逗号分隔
		"offer_sizes":                        this.Sell1Num,        // 卖一到卖五数量逗号分隔
		"offer_prices":                       this.Sell1Price,      // 卖一到卖五价格逗号分隔
		"base_stock_id":                      "",                   // 股票id
		"percent_change_from_previous_close": this.ChangePriceRate, // 涨跌幅百分比
		"change_from_previous_close":         this.ChangePrice,     // 涨跌幅
		"total_shares":                       "",                   // 总股本
		"non_restricted_shares":              "",                   // 非限制性股本
		"naps":                               "",                   // 每股净资产
		"eps":                                "",                   // 每股收益
		"amplitude":                          "",                   // 振幅
		"top_price":                          "",                   // 涨停价
		"bottom_price":                       "",                   // 跌停价
	}

	return &s
}
