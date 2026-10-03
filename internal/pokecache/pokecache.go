package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         sync.Mutex
	interval   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheEntry: map[string]cacheEntry{},
		interval:   interval,
	}
	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exist := c.cacheEntry[key]
	defer c.mu.Unlock()
	if exist {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	timer := time.NewTicker(c.interval)
	for range timer.C {
		c.mu.Lock()

		for key, _ := range c.cacheEntry {
			if time.Since(c.cacheEntry[key].createdAt) > c.interval {
				delete(c.cacheEntry, key)

			}
		}
		c.mu.Unlock()
	}

}
