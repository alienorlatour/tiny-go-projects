package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheRead(t *testing.T) {
	want := "ފަހެއް"
	cache := Cache[int, string]{
		data: map[int]string{5: want},
	}

	// Testing key is present.
	got, err := cache.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	// Testing key is absent.
	got, err = cache.Read(1)
	assert.NoError(t, err)
	assert.Equal(t, "", got)
}

func TestCacheUpsert(t *testing.T) {
	cache := Cache[int, string]{
		data: map[int]string{},
	}

	cache.Upsert(5, "fünf")
	assert.Equal(t, map[int]string{5: "fünf"}, cache.data)

	// Replace value for a present key.
	cache.Upsert(5, "pum")
	assert.Equal(t, map[int]string{5: "pum"}, cache.data)
}

func TestCacheDelete(t *testing.T) {
	cache := Cache[int, string]{
		data: map[int]string{6: "six"},
	}
	cache.Delete(6)
	assert.Equal(t, map[int]string{}, cache.data)
}
