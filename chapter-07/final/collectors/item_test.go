package collectors

type item string

// String is the default implementation of Stringer. Let's not overdo it.
func (i item) String() string {
	return string(i)
}

// Before is a simple implementation for our test type item.
func (i item) Before(s Sortable) bool {
	other, ok := s.(item)
	if !ok {
		return false
	}
	return i < other
}
