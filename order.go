package orderbook

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	UUID         string
	OrderType    OrderType
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	CurrencyPair CurrencyPair
	Time         time.Time
}

const (
	LimitBuyOrderType  orderType = "limit-buy-order"
	LimitSellOrderType orderType = "limit-sell-order"
)

func (o *Order) MatchesOrder(order *Order) bool {
	return o.OrderType.MatchesOrderType(order.OrderType) &&
		o.CurrencyPair.MatchesCurrencyPair(order.CurrencyPair) &&
		o.MatchesOrderPrice(order)
}

const orderBookFmtString = "|%20s|%12s|%12s|%32s|%38s|"

func (o Order) String() string {
	return fmt.Sprintf(orderBookFmtString,
		o.OrderType.Type(),
		o.Price.StringFixed(3),
		o.Quantity.StringFixed(8),
		o.Time.Format(time.RFC1123),
		o.UUID,
	)
}

type OrderMatcher interface {
	MatchesOrder(Order) bool
	OrderPriceMatcher
}

type OrderPriceMatcher interface {
	MatchesOrderPrice(OrderType) bool
}

func (o *Order) MatchesOrderPrice(order *Order) bool {
	buyOrderPrice := o.Price
	sellOrderPrice := order.Price
	if !o.OrderType.IsPriceSetter() {
		buyOrderPrice = order.Price
		sellOrderPrice = o.Price
	}
	return sellOrderPrice.LessThanOrEqual(buyOrderPrice)
}

type OrderType interface {
	fmt.Stringer
	OrderTypeMatcher
	Type() string

	// determines whether this order type sets
	// the price on a trade
	IsPriceSetter() bool
}

type orderType string

type OrderTypeMatcher interface {
	MatchesOrderType(OrderType) bool
}

func (t orderType) MatchesOrderType(orderType OrderType) bool {
	switch t {
	case LimitBuyOrderType:
		return orderType.Type() == string(LimitSellOrderType)
	case LimitSellOrderType:
		return orderType.Type() == string(LimitBuyOrderType)
	}
	// unknown order type
	return false
}

func (t orderType) IsPriceSetter() bool {
	return t == LimitBuyOrderType
}

func (t orderType) Type() string {
	return string(t)
}

func (t orderType) String() string {
	switch t {
	case LimitBuyOrderType:
		return "BUY"
	case LimitSellOrderType:
		return "SELL"
	}
	return ""
}
