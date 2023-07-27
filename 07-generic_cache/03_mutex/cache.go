package cache

import "sync"

// Cache is key-value storage.
type Cache[K comparable, V any] struct {
	mu   sync.RWMutex
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
func (c *Cache[K, V]) Read(key K) (V, bool) {
	// Lock the reading and the possible writing on the map
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, found := c.data[key]
	return v, found
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	// Lock the writing on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

// Delete removes the entry for the given key.
// If the key isn't present, Delete is a no-op.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
