package orderbook

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
)

type OrderbookTestSuite struct {
	suite.Suite
}

func TestOrderbookTestSuite(t *testing.T) {
	suite.Run(t, new(OrderbookTestSuite))
}

func (s *OrderbookTestSuite) TestNewOrderbook() {
	// Will cause a goroutine leak, todo in orderbook.go
	book := NewOrderbook()
	s.Require().IsType(&orderbook{}, book)
	impl := book.(*orderbook)
	s.Len(impl.orders, 0)
	s.Equal(bufferLength, cap(impl.tradeEmitter))
	s.Equal(bufferLength, cap(impl.orderRequestChan))
	time.Sleep(time.Millisecond * 100)
}

func (s *OrderbookTestSuite) TestAddOrder() {
	book := &orderbook{
		orderRequestChan: make(chan orderRequest, bufferLength),
	}
	go book.AddOrder(Order{UUID: uuid1})
	req := <-book.orderRequestChan
	s.Equal(uuid1, req.order.UUID)
	s.Equal(addOrderOperator, req.operator)
}

func (s *OrderbookTestSuite) TestRemoveOrder() {
	book := &orderbook{
		orderRequestChan: make(chan orderRequest, bufferLength),
	}
	go book.RemoveOrder(Order{UUID: uuid1})
	req := <-book.orderRequestChan
	s.Equal(uuid1, req.order.UUID)
	s.Equal(removeOrderOperator, req.operator)
}

func (s *OrderbookTestSuite) TestOrderRequestHandler() {
	bookInterface := NewOrderbook()
	book := bookInterface.(*orderbook)
	go func() {
		book.orderRequestChan <- orderRequest{
			operator: addOrderOperator,
			order: Order{
				OrderType: LimitBuyOrderType,
				UUID:      uuid1,
				Quantity:  one,
			},
		}
	}()
	time.Sleep(time.Millisecond * 10)
	s.Len(book.orders, 1)
	s.Equal(uuid1, book.orders[0].UUID)
}

func (s *OrderbookTestSuite) TestProcessIncomingOrderExecuteOnBuy() {
	bookInterface := NewOrderbook()
	book := bookInterface.(*orderbook)
	buyOrder := buildBuyOrder()
	book.orders = []*Order{
		buyOrder,
		buildSellOrder(),
	}

	book.processIncomingOrder(buyOrder)
	s.Len(book.orders, 1)
	s.Equal(uuid1, book.orders[0].UUID)
	s.True(one.Equal(book.orders[0].Quantity), book.orders[0].Quantity.String())
}

func (s *OrderbookTestSuite) TestProcessIncomingOrderExecuteOnSell() {
	bookInterface := NewOrderbook()
	book := bookInterface.(*orderbook)
	sellOrder := buildSellOrder()
	book.orders = []*Order{
		buildBuyOrder(),
		sellOrder,
	}

	book.processIncomingOrder(sellOrder)
	s.Len(book.orders, 1)
	s.Equal(uuid1, book.orders[0].UUID)
	s.True(one.Equal(book.orders[0].Quantity), book.orders[0].Quantity.String())
}

func (s *OrderbookTestSuite) TestAddOrderInternal() {
	bookInterface := NewOrderbook()
	book := bookInterface.(*orderbook)

	buyOrder100 := buildBuyOrder()
	buyOrder200 := buildBuyOrder()
	buyOrder300 := buildBuyOrder()
	buyOrder400 := buildBuyOrder()

	decimal100 := decimal.NewFromFloat(100)
	decimal200 := decimal.NewFromFloat(200)
	decimal300 := decimal.NewFromFloat(300)
	decimal400 := decimal.NewFromFloat(400)

	buyOrder100.Price = decimal100
	buyOrder200.Price = decimal200
	buyOrder300.Price = decimal300
	buyOrder400.Price = decimal400

	book.addOrder(buyOrder400)
	s.Len(book.orders, 1)
	book.addOrder(buyOrder300)
	s.Len(book.orders, 2)
	book.addOrder(buyOrder100)
	s.Len(book.orders, 3)

	s.Equal(decimal100, book.orders[0].Price)
	s.Equal(decimal300, book.orders[1].Price)
	s.Equal(decimal400, book.orders[2].Price)

	book.addOrder(buyOrder200)
	s.Len(book.orders, 4)

	s.Equal(decimal100, book.orders[0].Price)
	s.Equal(decimal200, book.orders[1].Price)
	s.Equal(decimal300, book.orders[2].Price)
	s.Equal(decimal400, book.orders[3].Price)
}
