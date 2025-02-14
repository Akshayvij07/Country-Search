package cache_test

import (
	"sync"
	"testing"

	"github.com/Akshayvij07/country-search/internals/cache"
)

// TestConcurrentCacheAccess tests the MapCache implementation for race conditions
// by concurrently writing and reading to the cache from multiple goroutines.
func TestConcurrentCacheAccess(t *testing.T) {
	c := cache.NewMapCache()
	var wg sync.WaitGroup
	key := "testKey"
	value := "testValue"

	// Set value in cache
	c.Set(key, value)

	// Spawn multiple goroutines to get and set values concurrently
	numRoutines := 100
	wg.Add(numRoutines * 2)

	for i := 0; i < numRoutines; i++ {
		go func() {
			defer wg.Done()
			_, _ = c.Get(key) // Read from cache
		}()

		go func(i int) {
			defer wg.Done()
			c.Set(key, value) // Write to cache
		}(i)
	}

	wg.Wait()
}
