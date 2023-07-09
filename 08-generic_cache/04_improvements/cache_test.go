package cache_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

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
	c := cache.New[int, string](3, time.Second)

	c.Upsert(5, "fünf")

	v, err := c.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "fünf", v)

	c.Upsert(5, "pum")

	v, err = c.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, err = c.Read(3)
	assert.NoError(t, err)
	assert.Equal(t, "drei", v)

	v, err = c.Read(5)
	assert.NoError(t, err)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	v, err = c.Read(5)
	assert.ErrorIs(t, err, cache.ErrNotFound)
	assert.Equal(t, "", v)

	v, err = c.Read(3)
	assert.NoError(t, err)
	assert.Equal(t, "drei", v)
}

// TestCache_Parallel_goroutines simulates a number of parallel tasks each operating on the cache.
// It passes if we only use "go test .", but we see the error as soon as we use "go test -race ."
func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string](5, time.Second)

	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			// Perform one operation that alters the content of the cache in each go routine.
			// The dataMutex prevents any race condition from happening.
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()
}

// TestCache_Parallel runs two routines that have concurrent access to write to the cache.
func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string](3, time.Second)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}

// TestCache_TTL inserts a value in a cache, and then makes sure we have reached its timeout.
// We expect an error to be returned in this case.
func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](5, time.Millisecond*100)
	c.Upsert("Norwegian", "Blue")
	time.Sleep(time.Millisecond * 200)

	// We've waited too long - the value's metabolic processes are now history.
	got, err := c.Read("Norwegian")

	assert.ErrorIs(t, err, cache.ErrExpired)
	assert.Equal(t, "", got)
}

func TestCache_MaxSize(t *testing.T) {
	t.Parallel()

	// Give it a TTL long enough to survive this test
	c := cache.New[int, int](3, time.Minute)

	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, err := c.Read(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, got)

	// Update 1, which will no longer make it the oldest
	c.Upsert(1, 10)

	// Adding a fourth element will discard the oldest - 2 in this case.
	c.Upsert(4, 4)

	// Trying to retrieve an element that should've been discarded by now.
	got, err = c.Read(2)
	assert.ErrorIs(t, err, cache.ErrNotFound)
	assert.Equal(t, 0, got)
}
