package money

// currency defines the code of a money and its precision.
type currency struct {
	code      string
	precision int
}

// errUnknownCurrency is returned when a currency is unsupported.
const errUnknownCurrency = moneyError("unknown currency")

// currencies defines the supported currencies.
var currencies = map[string]currency{
	"EUR": {code: "EUR", precision: 2},
	"USD": {code: "USD", precision: 2},
	"ETH": {code: "ETH", precision: 9},
}

// getCurrency returns the currency associated to a name and may return errUnknownCurrency.
func getCurrency(code string) (currency, error) {
	c, ok := currencies[code]
	if !ok {
		return currency{}, errUnknownCurrency
	}
	return c, nil
}
