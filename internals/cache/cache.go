package cache

import (
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/Akshayvij07/country-search/pkg/errors"
)

// Cache interface defining Get and Set methods
type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{})
}

// MapCache struct implementing Cache interface
type MapCache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewMapCache initializes a new cache instance
func NewMapCache() *MapCache {
	return &MapCache{
		data: make(map[string]interface{}),
	}
}

// Get retrieves a value from the cache
func (c *MapCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.data[key]
	if !found {
		return nil, errors.ErrKeyNotFound
	}
	return value, nil
}

// Set adds or updates a value in the cache
func (c *MapCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	c.printCache()
}

func (c *MapCache) printCache() {
	log.Info().Msgf("cache: %v", c.data)
}
