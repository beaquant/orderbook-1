package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zacatac/practice/rain/onsite/orderbook"
)

func main() {
	book := orderbook.NewOrderbook()

	go func(tradeEmitter <-chan orderbook.Trade) {
		for trade := range tradeEmitter {
			fmt.Println(trade)
		}
	}(book.TradeEmitter())

	filename := "orders.csv"
	orders, err := orderbook.OrdersFromCSV(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Placed orders")
	for _, order := range orders {
		fmt.Println(order)
	}
	fmt.Println()
	fmt.Println("Executed trades")
	for _, order := range orders {
		book.AddOrder(order)
	}

	// TODO(zackf): Signal that trades are still
	// processing, or leave process running waiting for
	// new orders.
	// Allow trades to process before exiting
	time.Sleep(time.Millisecond * 100)
	fmt.Println()
	fmt.Println("Final orderbook state")
	fmt.Println(book)

}
