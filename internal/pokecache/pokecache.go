package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) Cache {
	c := Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
	}

	go c.reapLoop(duration)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.cacheMap[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.reap(time.Now().UTC(), duration)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cacheMap {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheMap, k)
		}
	}
}
