package cache_test

import (
	"fmt"
	"sync"
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

	v, ok := c.Read(5)
	assert.True(t, ok)
	assert.Equal(t, "fünf", v)

	c.Upsert(5, "pum")

	v, ok = c.Read(5)
	assert.True(t, ok)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, ok = c.Read(3)
	assert.True(t, ok)
	assert.Equal(t, "drei", v)

	v, ok = c.Read(5)
	assert.True(t, ok)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	_, ok = c.Read(5)
	assert.False(t, ok)

	v, ok = c.Read(3)
	assert.True(t, ok)
	assert.Equal(t, "drei", v)
}

// TestCache_Parallel simulates a number of parallel tasks each operating on the cache.
// It passes if we only use "go test .", but we see the error as soon as we use "go test -race ."
func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string]()

	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			// Perform one operation that alters the content of the cache in each go routine.
			// The mutex prevents any race condition from happening.
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()
}
