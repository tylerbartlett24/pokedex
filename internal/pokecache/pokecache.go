package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]CacheEntry
	mu      *sync.Mutex
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = CacheEntry{
		createdAt: time.Now(),
		Val:       value,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.Entries[key]
	if !ok {
		return []byte{}, ok
	}
	return value.Val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}

type CacheEntry struct {
	createdAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		Entries: make(map[string]CacheEntry),
		mu:      &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
