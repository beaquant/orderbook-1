package orderbook

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"time"
)

type OrderTestSuite struct {
	suite.Suite
}

func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}

func (s *OrderTestSuite) TestMatchesOrderSuccess() {
	buyOrder, sellOrder := buildBuyOrder(), buildSellOrder()
	s.True(buyOrder.MatchesOrder(sellOrder))
	s.True(sellOrder.MatchesOrder(buyOrder))
}

func (s *OrderTestSuite) TestMatchesOrderFailure() {
	buyOrder, sellOrder := buildBuyOrder(), buildSellOrder()
	buyOrder.Price = sellOrder.Price.Sub(decimal.NewFromFloat(1.0))
	s.False(buyOrder.MatchesOrder(sellOrder))
	s.False(sellOrder.MatchesOrder(buyOrder))
}

func (s *OrderTestSuite) TestMatchesOrderType() {
	buyOrder, sellOrder := buildBuyOrder(), buildSellOrder()
	s.True(buyOrder.OrderType.MatchesOrderType(sellOrder.OrderType))
	s.True(sellOrder.OrderType.MatchesOrderType(buyOrder.OrderType))

	buyOrder.OrderType = LimitSellOrderType
	s.False(buyOrder.OrderType.MatchesOrderType(sellOrder.OrderType))
	s.False(sellOrder.OrderType.MatchesOrderType(buyOrder.OrderType))
}

func (s *OrderTestSuite) TestMatchesCurrencyPair() {
	buyOrder, sellOrder := buildBuyOrder(), buildSellOrder()
	s.True(buyOrder.CurrencyPair.MatchesCurrencyPair(sellOrder.CurrencyPair))
	s.True(sellOrder.CurrencyPair.MatchesCurrencyPair(buyOrder.CurrencyPair))

	buyOrder.CurrencyPair.CurrencyFrom = CurrencyUSD
	s.False(buyOrder.CurrencyPair.MatchesCurrencyPair(sellOrder.CurrencyPair))
	s.False(sellOrder.CurrencyPair.MatchesCurrencyPair(buyOrder.CurrencyPair))
}

func (s *OrderTestSuite) TestMatchesOrderPrice() {
	buyOrder, sellOrder := buildBuyOrder(), buildSellOrder()
	s.True(buyOrder.MatchesOrderPrice(sellOrder))
	s.True(sellOrder.MatchesOrderPrice(buyOrder))

	buyOrder.Price = sellOrder.Price.Sub(decimal.NewFromFloat(1.0))
	s.False(buyOrder.MatchesOrder(sellOrder))
	s.False(sellOrder.MatchesOrder(buyOrder))
}

func (s *OrderTestSuite) TestOrderString() {
	output := `|     limit-buy-order|     100.000|  1.00000000| Mon, 06 Aug 2018 19:26:26 +0000|  f360a27a-d255-46d9-8a7e-347d0b3fbd1e|`
	t, _ := time.Parse(time.RFC3339, "2018-08-06T19:26:26+00:00")
	order := Order{
		OrderType: LimitBuyOrderType,
		Quantity:  one,
		Price:     oneHundred,
		Time:      t,
		UUID:      "f360a27a-d255-46d9-8a7e-347d0b3fbd1e",
	}
	s.Equal(output, order.String())
}
