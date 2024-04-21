package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct {
	mu sync.Mutex
	cacheMap  map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()

		c.cacheMap[key] = cacheEntry{
			val: val,
			createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
 c.mu.Lock()
	defer c.mu.Unlock()
	
	entry, exists := c.cacheMap[key]
	if exists {
		return entry.val, true
	}else{
		return nil, false;
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

for range ticker.C {
	c.mu.Lock()
	for key, entry := range c.cacheMap {
		if time.Since(entry.createdAt) > interval {
			delete(c.cacheMap, key)
		}
	}
				c.mu.Unlock()
			}
		}


func NewCache(interval time.Duration)*Cache {
	
	    c := &Cache{
        cacheMap: make(map[string]cacheEntry),
    }
    go c.reapLoop(interval) // Start reaping on a goroutine
    return c 
}


