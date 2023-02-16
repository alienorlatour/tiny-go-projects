package ecbank

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ablqk/tiny-go-projects/chapter-05/5_bank/money"
)

type EuroCentralBank struct {
}

const path = ""

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (ecb EuroCentralBank) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrBadURL, path)
	}

	// use the default http client provided by the http library
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("http error:", err)
		return 0., fmt.Errorf("%w: %s", ErrServerSide, path)
	}

	fmt.Println(resp)

	return 0, nil
}
