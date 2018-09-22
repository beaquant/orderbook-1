package orderbook

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	uuid1 = "uuid-1"
	uuid2 = "uuid-2"
)

var (
	one        = decimal.NewFromFloat(1.0)
	two        = decimal.NewFromFloat(2.0)
	oneHundred = decimal.NewFromFloat(100.0)
	twoHundred = decimal.NewFromFloat(200.0)
)

func buildBuyOrder() *Order {
	return &Order{
		UUID:      uuid1,
		OrderType: LimitBuyOrderType,
		Price:     twoHundred,
		Quantity:  two,
		CurrencyPair: CurrencyPair{
			CurrencyFrom: CurrencyBTC,
			CurrencyTo:   CurrencyUSD,
		},
		Time: time.Now(),
	}
}

func buildSellOrder() *Order {
	return &Order{
		UUID:      uuid2,
		OrderType: LimitSellOrderType,
		Price:     oneHundred,
		Quantity:  one,
		CurrencyPair: CurrencyPair{
			CurrencyFrom: CurrencyBTC,
			CurrencyTo:   CurrencyUSD,
		},
		Time: time.Now(),
	}
}
