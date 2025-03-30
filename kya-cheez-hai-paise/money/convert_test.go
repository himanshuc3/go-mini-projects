package money_test

import (
	"kyacheezehaipaise/money"
	"kyacheezhaipaisa/money"
	"reflect"
	"testing"
)

type stubRate struct {
	rate money.ExchangeRate
	err  error
}

func (m stubRate) FetchExchangeRate(_, _ money.Currency) (money.ExchangeRate, error) {
	return m.rate, m.err
}

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to:     mustParseCurrency(t, "EUR"),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				expected := money.Amount{}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %v, got %v", expected, got)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)

	if err != nil {
		t.Fatalf("cannot parse currency %s code", code)
	}
	return currency
}

// NOTE:
// 1. Wrapper over ParseAmount to handle errors and deterministically
// return only value, not the error
func mustParseAmount(t *testing.T, value string, code string) money.Amount {
	// NOTE:
	// 1. Helper func - helps in debugging which test failed
	// by skipping the helper marked funcs and printing the line
	// number and stack trace from the test case
	t.Helper()

	n, err := money.ParseDecimal(value)

	if err != nil {
		t.Fatalf("invalid numbeR: %s", value)
	}

	currency, err := money.ParseCurrency(code)

	if err != nil {
		// NOTE:
		// 1. Fatalf - stops the test run immediately throwing the error
		t.Fatalf("invalid currency code: %s", code)
	}

	amount, err := money.NewAmount(n, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %v and currency code %s", n, code)
	}

	return amount
}
