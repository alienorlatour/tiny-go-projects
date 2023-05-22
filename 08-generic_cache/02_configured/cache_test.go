package cache_test

import (
	"github.com/stretchr/testify/assert"
	cache "learngo-pockets/genericcache"
	"testing"
)

/*
	TestCache is an integration test.

- Create a cache and check that it is empty?
- Upsert a new key <5, fünf> in the cache
- Read the value for this entry
- Upsert for the same entry with the new value
- Read the new value
- Upsert another key <3, drei> and check that is doesn't override
- Delete 5 and check that 3 still exists
*/
func TestCache(t *testing.T) {
	c := cache.New[int, string]()

	err := c.Upsert(5, "fünf")
	assert.NoError(t, err)

	v, err := c.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "fünf", v)

	err = c.Upsert(5, "pum")
	assert.NoError(t, err)

	v, err = c.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "pum", v)

	err = c.Upsert(3, "drei")
	assert.NoError(t, err)

	v, err = c.Read(3)
	assert.NoError(t, err)
	assert.Equal(t, "drei", v)

	v, err = c.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	_, err = c.Read(5)
	assert.ErrorIs(t, err, cache.ErrNotFound)

	v, err = c.Read(3)
	assert.NoError(t, err)
	assert.Equal(t, "drei", v)
}
