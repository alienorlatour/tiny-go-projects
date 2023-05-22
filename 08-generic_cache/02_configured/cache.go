package cache

import "sync"

// Cache is key, value storage.
type Cache[K comparable, V any] struct {
	data  map[K]V
	mutex sync.Mutex
}

// New create a new Cache with an initialised data.
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

// Read returns the associated value for a key,
// or ErrNotFound if the key is absent.
func (c *Cache[K, V]) Read(key K) (V, error) {
	v, ok := c.data[key]
	if !ok {
		return v, ErrNotFound
	}
	return v, nil
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	// Lock the writing on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value

	// Do not return an error for the moment,
	// but it can happen in the near future.
	return nil
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	delete(c.data, key)
}
