package cache

import (
	"slices"
	"sync"
	"time"
)

// Cache is key-value storage.
type Cache[K comparable, V any] struct {
	ttl time.Duration

	mu   sync.Mutex
	data map[K]entryWithExpiration[V]

	maxSize           int
	chronologicalKeys []K
}

type entryWithExpiration[V any] struct {
	value   V
	expires time.Time
}

// New create a new Cache with an initialised data.
func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]entryWithExpiration[V]),
		maxSize:           maxSize,
		chronologicalKeys: make([]K, 0, maxSize),
	}
}

// Read returns the associated value for a key,
// and a boolean to true if the key is absent.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	// Lock the reading and the possible writing on the cache.
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V

	e, ok := c.data[key]

	switch {
	case !ok:
		return zeroV, false
	case e.expires.Before(time.Now()):
		// The value has expired.
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return e.value, true
	}
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	// Lock the writing on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	// If the data is already present
	_, alreadyPresent := c.data[key]
	switch {
	case alreadyPresent:
		// Discard any previous reference.
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		// There is no room left in our map.
		// We need to remove an item to write a new one.
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	// Finally, insert the item.
	c.addKeyValue(key, value)
}

// Delete removes the entry for the given key.
// If the key isn't present, Delete is a no-op.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deleteKeyValue(key)
}

// addKeyValue inserts a key and its value into the cache.
func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = entryWithExpiration[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

// deleteKeyValue removes a key and its associated value from the cache.
func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronologicalKeys = slices.DeleteFunc(c.chronologicalKeys, func(k K) bool { return k == key })
	delete(c.data, key)
}
