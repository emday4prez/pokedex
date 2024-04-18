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



