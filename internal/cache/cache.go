package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mux      *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheMap: make(map[string]cacheEntry),
		mux:      &sync.Mutex{},
	}
	go cache.Clean(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) Clean(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.cacheMap {
			if time.Since(entry.createdAt) >= interval {
				delete(c.cacheMap, key)
			}
		}
		c.mux.Unlock()
	}
}
