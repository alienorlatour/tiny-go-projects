package money

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	quantity Quantity
	currency Currency
}
