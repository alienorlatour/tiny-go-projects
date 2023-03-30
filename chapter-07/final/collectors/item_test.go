package collectors

type item string

func (i item) Before(s Sortable) bool {
	other, ok := s.(item)
	if !ok {
		return false
	}
	return i < other
}
