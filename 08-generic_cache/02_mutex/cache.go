package cache

import "sync"

// Cache is key, value storage.
type Cache[K comparable, V any] struct {
	mutex sync.Mutex
	data  map[K]V
}

// New create a new Cache with an initialised data.
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

// Read returns the associated value for a key,
// or ErrNotFound if the key is absent.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	v, ok := c.data[key]
	if !ok {
		return v, false
	}
	return v, true
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	// Lock the writing on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}
