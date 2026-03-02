package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []Location
	next      string
	previous  string
}

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewLocCache(interval time.Duration) *Cache {
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
	for {
		time.Sleep(interval)
		c.mu.Lock()
		var to_delete []string
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				to_delete = append(to_delete, key)
			}
		}
		for _, key := range to_delete {
			delete(c.cache, key)
		}
		c.mu.Unlock()
	}
}

/////////////////////////////////////////////////////

type PokeCache struct {
	cache map[string]pokeCacheEntry
	interval time.Duration
	mu sync.Mutex
}

type pokeCacheEntry struct {
	createdAt time.Time
	pokemons []string
}

func NewPokeCache(interval time.Duration) *PokeCache {
	var new_cache PokeCache
	new_cache.cache = make(map[string]pokeCacheEntry)
	new_cache.interval = interval
	go new_cache.reapLoop(interval)
	return &new_cache
}

func (p *PokeCache) Add(location string, pokemons []string){
	p.mu.Lock()
	defer p.mu.Unlock()

	var entry pokeCacheEntry
	entry.pokemons = pokemons
	p.cache[location] = entry
}

func (p *PokeCache) Get(location string) ([]string, bool){
	p.mu.Lock()
	defer p.mu.Unlock()

	entry, ok := p.cache[location]
	if ok {
		return entry.pokemons, ok
	} else {
		return []string{}, ok
	}
	
}

func (c *PokeCache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		var to_delete []string
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				to_delete = append(to_delete, key)
			}
		}
		for _, key := range to_delete {
			delete(c.cache, key)
		}
		c.mu.Unlock()
	}
}