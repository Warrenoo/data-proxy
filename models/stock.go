package models

type Stock struct {
	Channel              string `redis:"-"`
	Id                   string `redis:"id,omitempty"`
	Symbol               string `redis:"symbol,omitempty"`
	Market               string `redis:"market,omitempty"`
	TradeTime            string `redis:"trade_time,omitempty"`
	RealTimePrice        string `redis:"last,omitempty"`
	ChangePrice          string `redis:"change_from_previous_close,omitempty"`
	ChangePriceRate      string `redis:"percent_change_from_previous_close,omitempty"`
	TodayHigh            string `redis:"high,omitempty"`
	TodayLow             string `redis:"low,omitempty"`
	Volume               string `redis:"volume,omitempty"`
	Buy1Price            string `redis:"bid_prices,omitempty"`
	Buy1Num              string `redis:"bid_sizes,omitempty"`
	Sell1Price           string `redis:"offer_prices,omitempty"`
	Sell1Num             string `redis:"offer_sizes,omitempty"`
	Week52High           string `redis:"high52_weeks,omitempty"`
	Week52Low            string `redis:"low52_weeks,omitempty"`
	Amount               string `redis:"total_value_trade,omitempty"`
	TodayOpen            string `redis:"open,omitempty"`
	LastClose            string `redis:"previous_close,omitempty"`
	Naps                 string `redis:"naps,omitempty"`
	Eps                  string `redis:"eps,omitempty"`
	Exchange             string `redis:"exchange,omitempty"`
	Amplitude            string `redis:"amplitude,omitempty"`
	TopPrice             string `redis:"top_price,omitempty"`
	BottomPrice          string `redis:"bottom_price,omitempty"`
	MarketCapitalization string `redis:"market_capitalization,omitempty"`
	PeRatio              string `redis:"pe_ratio,omitempty"`
}

//func (this *Stock) SaveFormat() *map[string]string {
//s := map[string]string{
//"high52_weeks":                       this.Week52High,
//"low52_weeks":                        this.Week52Low,
//"market_capitalization":              "",
//"pe_ratio":                           "",                   // 市盈率
//"exchange":                           "",                   // 交易所
//"symbol":                             this.Symbol,          // 股票代码
//"market":                             this.Market,          // 市场
//"trade_time":                         this.TradeTime,       // 交易时间
//"previous_close":                     this.LastClose,       // 昨收
//"open":                               this.TodayOpen,       // 今开
//"high":                               this.TodayHigh,       // 最高
//"low":                                this.TodayLow,        // 最低
//"last":                               this.RealTimePrice,   // 实时价格
//"volume":                             this.Volume,          // 成交量
//"total_value_trade":                  this.Amount,          // 成交额
//"bid_sizes":                          this.Buy1Num,         // 买一到买五数量逗号分隔
//"bid_prices":                         this.Buy1Price,       // 买一到买五价格逗号分隔
//"offer_sizes":                        this.Sell1Num,        // 卖一到卖五数量逗号分隔
//"offer_prices":                       this.Sell1Price,      // 卖一到卖五价格逗号分隔
//"base_stock_id":                      "",                   // 股票id
//"percent_change_from_previous_close": this.ChangePriceRate, // 涨跌幅百分比
//"change_from_previous_close":         this.ChangePrice,     // 涨跌幅
//"total_shares":                       "",                   // 总股本
//"non_restricted_shares":              "",                   // 非限制性股本
//"naps":                               "",                   // 每股净资产
//"eps":                                "",                   // 每股收益
//"amplitude":                          "",                   // 振幅
//"top_price":                          "",                   // 涨停价
//"bottom_price":                       "",                   // 跌停价
//}

//return &s
//}
