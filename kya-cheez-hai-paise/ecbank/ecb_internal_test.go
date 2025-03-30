package ecbank

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEuroCentralBank_FetchExchangeRate_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml...>`)
	}))

	defer ts.Close()

	ecb := Client{
		path: ts.URL,
	}

	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency("RON"))
}
