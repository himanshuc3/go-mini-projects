package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"
	"kyacheezhaipaisa/money"
	"os"
)

type rateError string

// NOTE:
// 1. Anything that implements this interface is an error
func (r rateError) Error() string {
	return string(r)
}

const baseCurrencyCode = "EUR"

const (
	ErrUnexpectedFormat   = rateError("unexpected format")
	ErrChangeRateNotFound = rateError("change Rate not found.")
)

type envelope struct {
	// NOTE:
	// 1. Can skip attributes in xml
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {

	decoder := xml.NewDecoder(respBody)

	var ecbMessage envelope

	err := decoder.Decode(&ecbMessage)

	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}

	rate, err := ecbMessage.exchangeRate(source, target)

	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}

	return rate, nil
}

func (e envelope) mappedExchangeRates() map[string]money.ExchangeRate {
	// NOTE:
	// 1. new() is used for returning a pointer
	// - Used for natives, structs, slices, maps, channels (zero-valued)
	// 2. make() - same same but diffelent
	// - returns initialized value of a specified type
	rates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	// NOTE:
	// 1. slices - Weird, but first value is index instead of value
	for _, c := range e.Rates {

		decimal, err := money.ParseDecimal(c.Rate)
		if err != nil {
			fmt.Printf("Error mapping exchange rates")
			os.Exit(1)
		}
		rates[c.Currency] = money.ExchangeRate(decimal)
	}
	euroDecimal, err := money.ParseDecimal("1.00")

	if err != nil {
		fmt.Printf("Error mapping exchange rates")
		os.Exit(1)
	}

	rates[baseCurrencyCode] = money.ExchangeRate(euroDecimal)
	return rates
}

func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		if val, err := money.ParseDecimal("1.00"); err != nil {
			return money.ExchangeRate{}, err
		} else {
			return money.ExchangeRate(val), nil
		}

	}

	rates := e.mappedExchangeRates()

	sourceFactor, sourceFound := rates[source]

	if !sourceFound {
		return money.ExchangeRate{}, fmt.Errorf("failed to find the source currency %s", source)
	}

	targetFactor, targetFound := rates[target]

	if !targetFound {
		return money.ExchangeRate{}, fmt.Errorf("failed to find target currency %s", target)
	}

	return targetFactor / sourceFactor
}
