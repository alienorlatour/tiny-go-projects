package cache

type cacheError string

func (e cacheError) Error() string {
	return string(e)
}

const (
	// ErrNotFound is returned when a key is not in the cache.
	ErrNotFound cacheError = "key not found"
	// ErrExpired is returned when a value in the cache has reached its TTL.
	ErrExpired cacheError = "value has reached TTL"
)
