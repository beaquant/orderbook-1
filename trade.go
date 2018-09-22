package orderbook

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Trade struct {
	OrderUUIDs []string
	Quantity   decimal.Decimal
	Price      decimal.Decimal
}

func NewTrade(order1, order2 Order) Trade {
	return Trade{
		OrderUUIDs: orderUUIDsByPriceSetter(order1, order2),
		Quantity:   tradeAmount(order1, order2),
		Price:      tradePrice(order1, order2),
	}
}

func tradeAmount(order1, order2 Order) decimal.Decimal {
	if order1.Quantity.LessThanOrEqual(order2.Quantity) {
		return order1.Quantity
	}
	return order2.Quantity
}

func tradePrice(order1, order2 Order) decimal.Decimal {
	if order1.OrderType.IsPriceSetter() {
		return order1.Price
	}
	return order2.Price
}

func orderUUIDsByPriceSetter(order1, order2 Order) []string {
	if order1.OrderType.IsPriceSetter() {
		return []string{order1.UUID, order2.UUID}
	}
	return []string{order2.UUID, order1.UUID}
}

func (t Trade) String() string {
	return fmt.Sprintf("Quantity: %s, Price: %s, Orders: %s",
		t.Quantity.StringFixed(8), t.Price.StringFixed(3), t.OrderUUIDs)
}
