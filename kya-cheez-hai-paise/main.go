package main

import (
	"flag"
	"fmt"
	"kyacheezhaipaisa/ecbank"
	"kyacheezhaipaisa/money"
	"os"
	"time"
)

// TODO:
// 1. Project not working because unclear about how
// to implement divide in Decimal format
// 2. Unit tests are pending

// NOTE:
// 1. Initiate a go module using
// go mod init module-name
// 2. If multiple binaries
// cmd/{binary-name}/main.go
// 3. Build the binary using
// go build -o convert main.go
// 4. Running from cmd
// go run . -from EUR -to CHF 6.69
// 5. flag.Args vs os.Args - flag leaves the mapped parameters
// 6. Compiled binary
// go build -o convert main.go
func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	flag.Parse()

	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s.\n", *from, err.Error())
		os.Exit(1)
	}

	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s.\n", *to, err.Error())
		os.Exit(1)
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s.\n", value, err.Error())
		os.Exit(1)
	}
	amount, err := money.NewAmount(quantity, fromCurrency)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	rates := ecbank.NewBank(30 * time.Second)

	convertedAmount, err := money.Convert(amount, toCurrency, rates)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert %s to %s: %s.\n", amount, toCurrency, err.Error())
		os.Exit(1)
	}

	fmt.Printf("Amount: %s; Currency: %s \n", amount, toCurrency)
	fmt.Printf("%s = %s\n", amount, convertedAmount)
}
