package pokecache

import (
	"time"
	"sync"
)

type CacheEntry struct {
	createdAt time.Time
	val 	  []byte
} // End CacheEntry struct

type Cache struct {
	cache map[string]CacheEntry
	mux   *sync.Mutex
} // End Cache struct

func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: make(map[string]CacheEntry),
		mux: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
} // End NewCache() func

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = CacheEntry{
		createdAt: time.Now().UTC(), 
		val: val,
	}
} // End Add(key, val) func

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val, ok :=  c.cache[key]
	return val.val, ok
} // End Get(key) func

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
} // End reapLoop() func

func (c *Cache) reap(currTime time.Time, prevTime time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(currTime.Add(-prevTime)) {
			delete(c.cache, k)
		}
	}
} // End reapLoop() func


// func main() {
// 	cache := NewCache()
// 	cache.Add("location", []byte{})
// 	cache.Add("Castle Rock", []byte{})
// 	cache.Add("Julesberg", []byte{})
// 	fmt.Println(cache.Get("CastleRock"))
// }


