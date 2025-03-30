package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// NOTE:
// (Question): How are we getting away with not defining the type here
const maxDecimal = 1e12

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")
	ErrTooLarge       = Error("quantity over 10^12 is too large")
)

type Decimal struct {
	subunits  int64
	precision byte
}

// NOTE:
// 1. Alias does not inherit the properties, because
// it doesn't have inheritance
// type A = B is different from type A B
type ExchangeRate Decimal

func ParseDecimal(value string) (Decimal, error) {
	// 1 - find the position of the . and split on it.
	// NOTE:
	// 1. Choose strings.cut over split

	intPart, fracPart, _ := strings.Cut(value, ".")

	// 2 - convert the string without the . to an integer. This could fail
	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	precision := byte(len(fracPart))

	// 4 - return the result
	return Decimal{subunits: subunits, precision: precision}, nil
}

func (d *Decimal) simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

func (d *Decimal) String() string {
	if d.precision == 0 {
		// NOTE:
		// 1. Question: Sprintf vs Printf ?
		return fmt.Sprintf("%d", d.subunits)
	}

	centsPerUnit := pow10(d.precision)
	frac := d.subunits % centsPerUnit
	integer := d.subunits / centsPerUnit

	decimalFormat := "%d.%" + strconv.Itoa(int(d.precision)) + "d"
	return fmt.Sprintf(decimalFormat, integer, frac)
}

func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000

	default:
		return int64(math.Pow(10, float64(power)))
	}
}
