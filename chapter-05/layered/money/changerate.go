package money

// ChangeRate is the rate to convert from a currency to another.
// It is a float32, because the precision of an official change rate is 5 significant digits.
type ChangeRate float32
