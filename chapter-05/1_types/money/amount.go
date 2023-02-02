package money

// Amount defines a quantity of money in a given currency.
type Amount struct {
	number   Number
	currency Currency
}
