package orderbook

import (
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
	"time"
)

const fixturesPath = "fixtures"
const ordersFixtureFilename = "orders.csv"

type CSVReaderTestSuite struct {
	suite.Suite
}

func TestCSVReaderTestSuite(t *testing.T) {
	suite.Run(t, new(CSVReaderTestSuite))
}

func (s *CSVReaderTestSuite) TestOrdersFromCSV() {
	orders, err := OrdersFromCSV(path.Join(fixturesPath, ordersFixtureFilename))
	s.NoError(err)
	s.Len(orders, 9)

	_, err = OrdersFromCSV("unknown-file")
	s.Error(err)
}

func (s *CSVReaderTestSuite) TestNewOrderFromCSV() {
	_, err := NewOrderFromCSV([]string{})
	s.Error(err)

	order, err := NewOrderFromCSV([]string{
		" buy ",
		"   1 ",
		" 100 ",
		" 2018-08-06T19:26:26+00:00 ",
	})
	s.NoError(err)
	s.Equal(LimitBuyOrderType, order.OrderType)
	s.True(one.Equal(order.Quantity))
	s.True(oneHundred.Equal(order.Price))
	t, _ := time.Parse(time.RFC3339, "2018-08-06T19:26:26+00:00")
	s.Equal(t, order.Time)

}

func (s *CSVReaderTestSuite) TestNewOrderTypeFromCSV() {
	orderType, err := NewOrderTypeFromCSV("buy")
	s.NoError(err)
	s.Equal(LimitBuyOrderType, orderType)

	orderType, err = NewOrderTypeFromCSV("sell")
	s.NoError(err)
	s.Equal(LimitSellOrderType, orderType)

	_, err = NewOrderTypeFromCSV("unknown-type")
	s.Error(err)

}
