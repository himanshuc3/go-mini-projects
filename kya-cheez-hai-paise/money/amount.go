package money

const (
	ErrTooPrecise = Error("quantity is too precise")
)

// NOTE:
// 1. Why choose custom type for quantity
// rather than float?
// - LLD: To create an interface, to prevent
// adding arbitrary precision
// - HLD: To prevent floating point errors
// 2. Testing with coverage
// go test ./... -cover
type Amount struct {
	quantity Decimal
	currency Currency
}

func (a Amount) String() string {
	return "[" + a.quantity.String() + " " + a.currency.code + "]"
}

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}
	quantity.subunits *= pow10(currency.precision - quantity.precision)
	quantity.precision = currency.precision
	return Amount{quantity: quantity, currency: currency}, nil
}

func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}
	return nil
}
