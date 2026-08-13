package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val		[]byte
}

type Cache struct {
	entries 	map[string]cacheEntry
	mu		sync.RWMutex
	interval	time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		entries:	make(map[string]cacheEntry),
		interval: 	interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
