package utils

import "sync"

// URLMap is safe to use concurrently.
type URLMap struct {
	mu sync.Mutex
	v  map[string]string
}

func (c *URLMap) Set(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] = key
	c.mu.Unlock()
}

func (c *URLMap) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func MakeMap() URLMap {
	return URLMap{v: make(map[string]string)}
}
