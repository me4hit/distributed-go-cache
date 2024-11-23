package cache

import (
	"container/list"
	"sync"
)

// LRUCache represents a thread-safe LRU cache
type LRUCache struct {
	capacity int
	mu       sync.Mutex
	cache    map[string]*list.Element
	list     *list.List
}

// entry represents a key-value pair in the cache
type entry struct {
	key   string
	value string
}

// NewLRUCache creates a new LRUCache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return "", false
}

// Set adds a key-value pair to the cache
func (c *LRUCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if c.list.Len() >= c.capacity {
		// Evict least recently used
		back := c.list.Back()
		if back != nil {
			c.list.Remove(back)
			delete(c.cache, back.Value.(*entry).key)
		}
	}

	e := &entry{key: key, value: value}
	elem := c.list.PushFront(e)
	c.cache[key] = elem
}
