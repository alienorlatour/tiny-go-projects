package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheRead(t *testing.T) {
	// Initialise the cache and prepare data.
	cache := New[int, string](5, time.Second)
	cache.data[5] = entryWithTimeout[string]{value: "ފަހެއް", expires: time.Now().Add(time.Minute)}

	// Testing key is present.
	got, err := cache.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "ފަހެއް", got)

	// Testing key is absent.
	got, err = cache.Read(1)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Equal(t, "", got)
}

func TestCacheUpsert(t *testing.T) {
	cache := New[int, string](5, time.Second)

	cache.Upsert(5, "fünf")
	assert.Equal(t, "fünf", cache.data[5].value)

	// Replace value for a present key.
	cache.Upsert(5, "pum")
	assert.Equal(t, "pum", cache.data[5].value)
}

func TestCacheDelete(t *testing.T) {
	cache := New[int, string](5, time.Second)
	cache.Upsert(6, "six")

	cache.Delete(6)
	assert.Equal(t, map[int]entryWithTimeout[string]{}, cache.data)
}
