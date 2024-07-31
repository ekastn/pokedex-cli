package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mut   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mut:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	item := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = item
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(time.Now().Add(-interval)) {
			delete(c.cache, k)
		}
	}
}
