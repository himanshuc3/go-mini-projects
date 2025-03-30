package money

import (
	"fmt"
	"math"
)

func Convert(amount Amount, to Currency, rates exchangeRatesInterface) (Amount, error) {
	r, err := rates.FetchExchangeRate(amount.currency, to)

	fmt.Println(r)

	if err != nil {

		return Amount{}, fmt.Errorf("cannot get change rate: %w", err)
	}

	convertedValue := applyExchangeRate(amount, to, r)

	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}
	return convertedValue, nil
}

func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted, err := multiply(a.quantity, rate)
	if err != nil {
		return Amount{}
	}

	precisionDiff := byte(math.Abs(float64(converted.precision) - float64(target.precision)))

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(precisionDiff)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(precisionDiff)
	}
	converted.precision = target.precision
	return Amount{
		currency: target,
		quantity: converted,
	}
}

func multiply(d Decimal, rate ExchangeRate) (Decimal, error) {

	dec := Decimal{
		subunits:  d.subunits * rate.subunits,
		precision: d.precision + rate.precision,
	}
	dec.simplify()
	return dec, nil
}

func divide(d Decimal, rate ExchangeRate) (Decimal, error) {
	dec := Decimal{
		subunits:  d.subunits * rate.subunits,
		precision: d.precision + rate.precision,
	}
	dec.simplify()
	return dec, nil
}
