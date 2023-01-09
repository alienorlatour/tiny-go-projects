package money

const (
	// ErrInputTooSmall is returned if the amount to convert is too small, which could lead to precision issues.
	ErrInputTooSmall = moneyError("input amount should be at least 1.00")
	// ErrOutputTooSmall is returned if the converted amount is too small, in order to protect against precision issues.
	ErrOutputTooSmall = moneyError("output amount is less than 1.00")
	// ErrInputTooLarge is returned if the input amount is too large - this would cause floating point precision errors.
	ErrInputTooLarge = moneyError("input amount over 10^15 is too large")
	// ErrOutputTooLarge is returned if the converted amount is too large, to protect against floating point errors.
	ErrOutputTooLarge = moneyError("output amount is too large (over 10^15)")

	minAmount = 1
	// maxAmount value is a thousand billion, using the short scale -- 10^12.
	maxAmount = 1_000_000_000_000
)

func (n number) validateInput() error {
	switch {
	case n.tooSmall():
		return ErrInputTooSmall
	case n.tooBig():
		return ErrInputTooLarge
	}
	return nil
}

func (n number) validateOutput() error {
	switch {
	case n.tooSmall():
		return ErrOutputTooSmall
	case n.tooBig():
		return ErrOutputTooLarge
	}
	return nil
}

func (n number) tooSmall() bool {
	return n.integerPart < minAmount
}

func (n number) tooBig() bool {
	return n.integerPart > maxAmount
}
