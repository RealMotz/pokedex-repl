package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type PokeCache struct {
	mu       *sync.Mutex
	cache    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) PokeCache {
	pokeCache := PokeCache{
		mu:       &sync.Mutex{},
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}

	go pokeCache.reapLoop()
	return pokeCache
}

func (c *PokeCache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cache[key]; !ok {
		return []byte{}, false
	}

	val := c.cache[key].val
	return val, true
}

func (c *PokeCache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	<-ticker.C
	for key, val := range c.cache {
		now := time.Now()
		diff := now.Sub(val.createdAt)
		if diff > c.interval {
			c.mu.Lock()
			delete(c.cache, key)
			c.mu.Unlock()
		}
	}
}
