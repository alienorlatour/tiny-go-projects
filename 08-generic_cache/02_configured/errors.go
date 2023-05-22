package cache

type cacheError string

func (e cacheError) Error() string {
	return string(e)
}

const ErrNotFound cacheError = "key not found"
