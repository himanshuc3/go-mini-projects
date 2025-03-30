package ecbank

import (
	"fmt"
	"kyacheezhaipaisa/money"
	"net/http"
	"time"
)

const (
	ErrServerSide        = ecbankError("error calling server")
	ErrClientSide        = ecbankError("error calling server")
	ErrUnknownStatuscode = ecbankError("unknown error with exchange rate")
)

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}

func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		// NOTE:
		// 1. Errorf wraps the original error
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatuscode, statusCode)
	}
}

// NOTE:
// 1. Client for our application
type Client struct {
	client http.Client
}

func NewBank(timeout time.Duration) Client {
	return Client{
		client: http.Client{Timeout: timeout},
	}
}

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {

	const path = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := c.client.Get(path)

	fmt.Println(err)

	if err != nil {
		// NOTE:
		// 1. We know the error returned is of this type
		// eurlErr, ok := err.(*url.Error)
		// 2. Alternative idiomatic way
		// var urlErr *url.Error
		// if ok := errors.As(err, &urlErr); ok && urlErr.Timeout() {

		// }

		// NOTE:
		// 1. %w is for formatting errors
		return money.ExchangeRate{}, fmt.Errorf("%w, %s", ErrServerSide, err.Error())
	}
	defer resp.Body.Close()

	rate, err := readRateFromResponse(source.ISOCode(), target.ISOCode(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}
	return rate, nil
}
