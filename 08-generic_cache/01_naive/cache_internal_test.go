package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheRead(t *testing.T) {
	want := "ފަހެއް"
	cache := Cache[int, string]{
		map[int]string{5: want},
	}

	// Testing key is present.
	got, ok := cache.Read(5)
	assert.True(t, ok)
	assert.Equal(t, want, got)

	// Testing key is absent.
	got, ok = cache.Read(1)
	assert.False(t, ok)
	assert.Equal(t, "", got)
}

func TestCacheUpsert(t *testing.T) {
	cache := Cache[int, string]{
		map[int]string{},
	}

	cache.Upsert(5, "fünf")
	assert.Equal(t, map[int]string{5: "fünf"}, cache.data)

	// Replace value for a present key.
	cache.Upsert(5, "pum")
	assert.Equal(t, map[int]string{5: "pum"}, cache.data)
}

func TestCacheClear(t *testing.T) {
	cache := Cache[int, string]{
		map[int]string{6: "six"},
	}
	cache.Clear(6)
	assert.Equal(t, map[int]string{}, cache.data)
}
