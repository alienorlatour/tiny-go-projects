package cache

import (
	"sync"
	"time"
)

// Cache is key, value storage.
type Cache[K comparable, V any] struct {
	now     func() time.Time
	refresh time.Duration

	mutex sync.Mutex
	data  map[K]*entry[V]
}

type entry[V any] struct {
	value   V
	expires time.Time
}

// New create a new Cache with an initialised data.
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		now: func() time.Time {
			return time.Now()
		},
		refresh: time.Minute, // We might want to configure it.

		data: make(map[K]*entry[V]),
	}
}

// Read returns the associated value for a key,
// or ErrNotFound if the key is absent.
func (c *Cache[K, V]) Read(key K, load func(K) (V, error)) (V, error) {
	// Lock the reading and the possible writing on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()

	e, ok := c.data[key]
	if ok && e.expires.After(c.now()) {
		// If the value exists and its TTL is expired,
		// let's deal with it.
		ok = false
	}

	// If the value does not exist or its TTL is expired...
	if !ok {
		// load the value...
		value, err := load(key)
		if err != nil {
			return value, err
		}

		// update its expires time.
		e = &entry[V]{
			value:   value,
			expires: c.now().Add(c.refresh),
		}
	}

	return e.value, nil
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	// Lock the writing on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Insert the value if it does not exist,
	// otherwise update it.
	c.data[key] = &entry[V]{
		value:   value,
		expires: c.now().Add(c.refresh),
	}
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}
