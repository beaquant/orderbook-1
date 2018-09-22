package orderbook

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TradeTestSuite struct {
	suite.Suite
}

func TestTradeTestSuite(t *testing.T) {
	suite.Run(t, new(TradeTestSuite))
}

func (s *TradeTestSuite) TestNewTrade() {
	buyOrder, sellOrder := *buildBuyOrder(), *buildSellOrder()
	trade := NewTrade(buyOrder, sellOrder)
	s.Equal(twoHundred, trade.Price)
	s.Equal(one, trade.Quantity)
	s.Require().Len(trade.OrderUUIDs, 2)
	s.Equal(uuid1, trade.OrderUUIDs[0])
	s.Equal(uuid2, trade.OrderUUIDs[1])

	tradeTwo := NewTrade(sellOrder, buyOrder)
	s.Equal(trade, tradeTwo)
}

func (s *TradeTestSuite) TestTradeString() {
	t := Trade{
		OrderUUIDs: []string{"1", "2"},
		Quantity:   one,
		Price:      oneHundred,
	}
	s.Equal("Quantity: 1.00000000, Price: 100.000, Orders: [1 2]", t.String())
}
