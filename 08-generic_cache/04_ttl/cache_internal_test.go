package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheRead(t *testing.T) {
	// Initialise the cache and prepare data.
	cache := New[int, string]()
	cache.data[5] = &entry[string]{value: "ފަހެއް", expires: time.Now().Add(time.Minute)}

	// Testing key is present.
	got, err := cache.Read(5, func(i int) (string, error) {
		return cache.data[5].value, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, "ފަހެއް", got)

	// Testing key is absent.
	got, err = cache.Read(1, func(i int) (string, error) {
		return "", ErrNotFound
	})
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Equal(t, "", got)
}

func TestCacheUpsert(t *testing.T) {
	cache := New[int, string]()

	cache.Upsert(5, "fünf")
	assert.Equal(t, "fünf", cache.data[5].value)

	// Replace value for a present key.
	cache.Upsert(5, "pum")
	assert.Equal(t, "pum", cache.data[5].value)
}

func TestCacheClear(t *testing.T) {
	cache := New[int, string]()
	cache.data[6] = &entry[string]{value: "six", expires: time.Now().Add(time.Minute)}

	cache.Clear(6)
	assert.Equal(t, map[int]*entry[string]{}, cache.data)
}
