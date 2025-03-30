package money

type exchangeRatesInterface interface {
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}
