package cache

// Cache is key, value storage.
type Cache[K comparable, V any] struct {
	data map[K]V
}

// New create a new Cache with an initialised data.
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

// Read returns the associated value for a key,
// and a boolean to true if the key is absent.
func (c *Cache[K, V]) Read(key K) (V, error) {
	v, ok := c.data[key]
	if !ok {
		return v, ErrNotFound
	}

	return v, nil
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	c.data[key] = value
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}
