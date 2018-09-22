package orderbook

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Orderbook interface {
	fmt.Stringer

	AddOrder(Order)
	RemoveOrder(Order)
	TradeEmitter() <-chan Trade
}

type orderbook struct {
	orders []*Order

	// resulting trades are emitted exactly once
	// on this channel
	tradeEmitter chan Trade

	// orderRequestChan is a channel which serializes
	// all incoming requests to edit the orderbook
	orderRequestChan chan orderRequest
}

const bufferLength = 10000

func NewOrderbook() Orderbook {
	book := &orderbook{
		orders:           []*Order{},
		tradeEmitter:     make(chan Trade, bufferLength),
		orderRequestChan: make(chan orderRequest, bufferLength),
	}
	go book.orderRequestHandler()
	return book
}

func (b *orderbook) AddOrder(order Order) {
	// TODO(zacatac): incoming order validation
	b.orderRequestChan <- orderRequest{
		operator: addOrderOperator,
		order:    order,
	}
}

func (b *orderbook) TradeEmitter() <-chan Trade {
	return b.tradeEmitter
}

func (b *orderbook) RemoveOrder(order Order) {
	// TODO(zacatac): incoming order validation
	b.orderRequestChan <- orderRequest{
		operator: removeOrderOperator,
		order:    order,
	}
}

type orderRequest struct {
	operator operator
	order    Order
}

type operator string

const (
	addOrderOperator    operator = "add"
	removeOrderOperator operator = "remove"
)

func (b *orderbook) orderRequestHandler() {
	// ensuring that requests are serialized
	// TODO(zacatac): Add exit channel to cleanly exit
	for orderRequest := range b.orderRequestChan {
		order := orderRequest.order
		switch orderRequest.operator {
		case addOrderOperator:
			b.addOrder(&order)
			// will potentially block for a longer period,
			// but user input is not blocked
			b.processIncomingOrder(&order)
		case removeOrderOperator:
			b.removeOrder(&order)
		}
	}
}

func (b *orderbook) processIncomingOrder(incomingOrder *Order) {
	lenOrders := len(b.orders)

	for i := 0; i < lenOrders; i++ {
		order := b.orders[i]

		if !incomingOrder.MatchesOrder(order) {
			continue
		}

		trade := NewTrade(*incomingOrder, *order)

		if trade.Quantity.IsZero() {
			continue
		}

		// update quantities in place
		incomingOrder.Quantity = incomingOrder.Quantity.Sub(trade.Quantity)
		order.Quantity = order.Quantity.Sub(trade.Quantity)

		// decide whether order has been entirely filled
		if order.Quantity.IsZero() {
			b.removeOrderAtIndex(i)
			lenOrders--
		}

		// emit trade
		b.tradeEmitter <- trade
	}
	// Look through the orders again to remove
	// the incoming order from the list
	if incomingOrder.Quantity.IsZero() {
		for i, order := range b.orders {
			if order.UUID == incomingOrder.UUID {
				b.removeOrderAtIndex(i)
			}
		}
	}
}

func (b *orderbook) removeOrderAtIndex(i int) {
	if i >= len(b.orders) {
		return
	}
	if i == len(b.orders)-1 {
		b.orders = b.orders[:i]
	} else {
		// guaranteed to still be sorted
		b.orders = append(b.orders[:i], b.orders[i+1:]...)
	}
}

// addOrder performs a binary search to find where
// to insert an incoming order
func (b *orderbook) addOrder(order *Order) {
	if len(b.orders) == 0 {
		b.orders = []*Order{order}
		return
	}
	i := search(b.orders, order, 0, len(b.orders)-1)
	b.orders = append(b.orders, nil)
	copy(b.orders[i+1:], b.orders[i:])
	b.orders[i] = order
}

func search(orders []*Order, order *Order, left, right int) int {
	if right <= left {
		if order.Price.LessThan(orders[left].Price) {
			return left
		}
		return left + 1
	}

	mid := (left + right) / 2

	midOrderPrice := orders[mid].Price
	if order.Price.Equal(midOrderPrice) {
		return mid + 1
	}

	if order.Price.LessThan(midOrderPrice) {
		right = mid - 1
	} else {
		left = mid + 1
	}

	return search(orders, order, left, right)
}

func (b *orderbook) removeOrder(orderToRemove *Order) error {
	for i, order := range b.orders {
		if order.UUID != orderToRemove.UUID {
			continue
		}
		if i == len(b.orders)-1 {
			b.orders = b.orders[:i]
		} else {
			// guaranteed to still be sorted
			b.orders = append(b.orders[:i], b.orders[i+1:]...)
		}
		return nil
	}
	return errors.New("RemoveOrderError: order not found")
}

func (b *orderbook) String() string {
	table := fmt.Sprintf(orderBookFmtString, "Type", "Price", "Quantity", "Time", "UUID") + "\n"

	for _, order := range b.orders {
		if order.OrderType.Type() == string(LimitBuyOrderType) {
			table += fmt.Sprintln(order)
		}
	}

	table += fmt.Sprintf(orderBookFmtString, "", "", "", "", "") + "\n"

	for _, order := range b.orders {
		if order.OrderType.Type() == string(LimitSellOrderType) {
			table += fmt.Sprintln(order)
		}
	}

	return table
}

func (b *orderbook) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.orders)
}
