package cache

import (
	"slices"
	"testing"
	"time"

	"golang.org/x/exp/maps"
)

// TestCache is an internal test that ensures our cache updates synchronously its internal data
func TestCache(t *testing.T) {
	c := New[int, int](3, time.Second)

	if c.data == nil {
		t.Error("Data content not initialised")
	}
	if c.chronologicalKeys == nil {
		t.Error("List of keys not initialised")
	}

	c.Upsert(1, 10)
	c.Upsert(2, 20)
	c.Upsert(3, 30)

	expectedKeys := []int{1, 2, 3}
	if !slices.Equal(c.chronologicalKeys, expectedKeys) {
		t.Errorf("List of keys should be %v, instead we have %v", expectedKeys, c.chronologicalKeys)
	}

	dataKeys := maps.Keys(c.data)
	// This is a slice of its own, we can sort it.
	slices.Sort(dataKeys)
	if !slices.Equal(dataKeys, expectedKeys) {
		t.Errorf("Keys of the map should be %v, instead we have %v", expectedKeys, dataKeys)
	}

	c.Upsert(2, 31)
	expectedKeys = []int{1, 3, 2}
	if !slices.Equal(c.chronologicalKeys, expectedKeys) {
		t.Errorf("After upserting: list of keys should be %v, instead we have %v", expectedKeys, c.chronologicalKeys)
	}

	c.Delete(3)
	dataKeys = maps.Keys(c.data)
	// This is a slice of its own, we can sort it.
	slices.Sort(dataKeys)
	expectedKeys = []int{1, 2}
	if !slices.Equal(dataKeys, expectedKeys) {
		t.Errorf("After deleting a value: keys of the map should be %v, instead we have %v", expectedKeys, dataKeys)
	}

	value, found := c.Read(1)
	if !found {
		t.Error("Key 1 should be found")
	}
	if value != 10 {
		t.Error("Value of key 1 should be 10")
	}

	// Reach TTL
	time.Sleep(time.Second)

	value, found = c.Read(1)
	if found {
		t.Error("Key 1 should have been discarded")
	}
	if value != 0 {
		t.Errorf("Value of unfound key 1 shouuld be 0, got %d", value)
	}
}
