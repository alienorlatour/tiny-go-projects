package money

func NewAmountHelper(value string, targetCurrency string) Amount {
	number, _ := ParseNumber(value)
	currency, _ := ParseCurrency(targetCurrency)

	return Amount{number: number, currency: currency}
}
