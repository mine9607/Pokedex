package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	// my code
	mu sync.RWMutex

	// solution code
	// mux *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMap: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval) // start the reaper in a goroutine
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// function that adds a new entry to the cache
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheMap[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	cache := c.cacheMap
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range cache {
		if val.createdAt.Before(now.Add(-last)) {
			delete(cache, key)
		}
	}
}
