package orderbook

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func OrdersFromCSV(filename string) (orders []Order, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	var record []string
	for {
		record, err = csvReader.Read()
		if err != nil {
			// signals the end of the file
			if err == io.EOF {
				return orders, nil
			}
			return nil, err
		}

		order, orderParseErr := NewOrderFromCSV(record)
		if orderParseErr != nil {
			return nil, orderParseErr
		}
		orders = append(orders, order)
	}
}

// example CSV line:
// buy, 1.00000000, 100.00, 2018-08-06T19:26:26+00:00
func NewOrderFromCSV(record []string) (Order, error) {

	// TODO(zacatac): import uuid gen library
	if len(record) != 4 {
		return Order{}, errors.New("invalid number of CSV columns")
	}

	for i := range record {
		record[i] = strings.TrimSpace(record[i])
	}

	orderUUID, err := uuid.NewV4()
	if err != nil {
		return Order{}, err
	}

	orderType, err := NewOrderTypeFromCSV(record[0])
	if err != nil {
		return Order{}, err
	}

	quantity, err := decimal.NewFromString(record[1])
	if err != nil {
		return Order{}, err
	}

	price, err := decimal.NewFromString(record[2])
	if err != nil {
		return Order{}, err
	}

	t, err := time.Parse(time.RFC3339, record[3])
	if err != nil {
		return Order{}, err
	}

	return Order{
		UUID:      orderUUID.String(),
		OrderType: orderType,
		Price:     price,
		Quantity:  quantity,
		Time:      t,
		// TODO(zacatac): Support other currency pairs
		CurrencyPair: CurrencyPair{
			CurrencyFrom: CurrencyUSD,
			CurrencyTo:   CurrencyBTC,
		},
	}, nil
}

func NewOrderTypeFromCSV(orderTypeStr string) (OrderType, error) {
	switch orderTypeStr {
	case "buy":
		return LimitBuyOrderType, nil
	case "sell":
		return LimitSellOrderType, nil
	}
	return nil, errors.New("invalid order type")
}
