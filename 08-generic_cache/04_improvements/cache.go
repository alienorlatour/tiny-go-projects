package cache

import (
	"sync"
	"time"
)

// Cache is key, value storage.
type Cache[K comparable, V any] struct {
	ttl time.Duration

	dataMutex sync.RWMutex
	data      map[K]*entryWithTimeout[V]

	maxSize    int
	oldestKeys []K
}

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time
}

// New create a new Cache with an initialised data.
func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:        ttl,
		data:       make(map[K]*entryWithTimeout[V]),
		maxSize:    maxSize,
		oldestKeys: make([]K, 0, maxSize),
	}
}

// Read returns the associated value for a key,
// or ErrNotFound if the key is absent.
func (c *Cache[K, V]) Read(key K) (V, error) {
	// Lock the reading and the possible writing on the cache.
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()

	var zeroV V

	e, ok := c.data[key]

	switch {
	case !ok:
		// e is the zero-value, e.value is also the zero-value.
		return zeroV, ErrNotFound
	case e.expires.Before(time.Now()):
		// The value has expired.
		c.deleteKeyValue(key)
		return zeroV, ErrExpired
	default:
		return e.value, nil
	}
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	// Lock the writing on the map
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()

	// If the data is already present
	_, alreadyPresent := c.data[key]
	switch {
	case alreadyPresent:
		// Mark key as the last inserted.
		c.deleteKeyValue(key)
		c.addKeyValue(key, value)
	case len(c.data) == c.maxSize:
		// There is no room left in our map.
		// We need to remove an item to write a new one.
		c.deleteOldest()
		c.addKeyValue(key, value)
	default:
		// Inserting a new item - no special action to perform.
		c.addKeyValue(key, value)
	}
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()

	c.deleteKeyValue(key)
}

func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = &entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.oldestKeys = append(c.oldestKeys, key)
}

func (c *Cache[K, V]) deleteKeyValue(key K) {
	// Find the key in the list.
	for index, k := range c.oldestKeys {
		if k == key {
			// We've found the correct index.
			// Delete the element from the slice...
			c.oldestKeys = append(c.oldestKeys[:index], c.oldestKeys[index+1:]...)
			delete(c.data, key)
			return
		}
	}
}

// deleteOldest removes the oldest item in our map.
func (c *Cache[K, V]) deleteOldest() {
	if len(c.oldestKeys) == 0 {
		return
	}

	// Discard the most ancient element.
	c.deleteKeyValue(c.oldestKeys[0])
}
