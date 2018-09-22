package orderbook

type CurrencyPair struct {
	CurrencyFrom Currency
	CurrencyTo   Currency
}

type Currency string

const (
	CurrencyBTC Currency = "BTC"
	CurrencyUSD Currency = "USD"
)

func (c *CurrencyPair) MatchesCurrencyPair(pair CurrencyPair) bool {
	return c.CurrencyFrom == pair.CurrencyFrom &&
		c.CurrencyTo == pair.CurrencyTo
}
