package money

import "regexp"

const ErrInvalidCurrencyCode = Error("invalid currency code")

type Currency struct {
	code      string
	precision byte
}

// NOTE:
// 1. Stringer interface implemented
// - fmt will use this to print the struct
func (c Currency) String() string {
	return c.code
}

func (c Currency) ISOCode() string {
	return c.code
}

func ParseCurrency(code string) (Currency, error) {

	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	match, _ := regexp.MatchString("[A-Z]{3}", code)

	if !match {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}
