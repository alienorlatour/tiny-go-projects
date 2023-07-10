package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cache "learngo-pockets/genericcache"
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

	c.Upsert(5, "fünf")

	v, found := c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "fünf", v)

	c.Upsert(5, "pum")

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	v, found = c.Read(5)
	assert.False(t, found)
	assert.Equal(t, "", v)

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)
}
