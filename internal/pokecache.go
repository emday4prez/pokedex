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

func NewCache(interval time.Duration)Cache {
	
	return Cache{}
}


