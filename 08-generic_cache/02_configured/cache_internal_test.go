package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheRead(t *testing.T) {
	want := "coucou"
	cache := Cache[int, string]{
		data: map[int]string{5: want},
	}

	// Testing key is present.
	got, err := cache.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	// Testing key is absent.
	got, err = cache.Read(1)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Equal(t, "", got)
}

func TestCacheUpsert(t *testing.T) {
	cache := Cache[int, string]{
		data: map[int]string{},
	}

	err := cache.Upsert(5, "fünf")
	assert.NoError(t, err)
	assert.Equal(t, map[int]string{5: "fünf"}, cache.data)

	// Replace value for a present key.
	err = cache.Upsert(5, "pum")
	assert.NoError(t, err)
	assert.Equal(t, map[int]string{5: "pum"}, cache.data)
}

func TestCacheDelete(t *testing.T) {
	cache := Cache[int, string]{
		data: map[int]string{6: "six"},
	}
	cache.Delete(6)
	assert.Equal(t, map[int]string{}, cache.data)
}
