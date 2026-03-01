package internal

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []Location
	next string
	previous string
}

type Cache struct {
	cache map[string]cacheEntry
	interval time.Duration
	mu sync.Mutex
}

func NewCache(interval time.Duration)(*Cache) {
	var new_cache Cache
	new_cache.cache = make(map[string]cacheEntry)
	go new_cache.reapLoop(interval)
	return &new_cache
}

func (c *Cache) Add(key string, val []Location, next string, previous string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var new_entry cacheEntry
	new_entry.val = val
	new_entry.next = next
	new_entry.previous = previous
	c.cache[key] = new_entry
}

func (c *Cache) Get(key string) (cacheEntry, bool) {
	var emptycache cacheEntry
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if ok {
		return val, ok
	} else {
		return emptycache, ok
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	for ;; {
		time.Sleep(interval)
		c.mu.Lock()
		var to_delete []string
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval{
				to_delete = append(to_delete, key)
			}
		}
		for _, key := range to_delete {
			delete(c.cache, key)
		}
		c.mu.Unlock()
	}
}