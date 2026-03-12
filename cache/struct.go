package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
}

type CacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache() Cache {
	c := Cache {
		cache: make(map[string]CacheEntry),
	}

	// call the cleanRoutine in sepearte thread
	// go c.cleanRoutine(1)

	return c
}


func (c *Cache) Add(key string, entry []byte) {
	c.cache[key] = CacheEntry{
		val: entry,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cache, ok := c.cache[key]
	return cache.val, ok
}

// launch teh cleanup function in ago routine which will run at defined interval
func (c *Cache) cleanRoutine(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.cleanup(interval)
	}
}

func (c *Cache) cleanup(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	for k,v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
