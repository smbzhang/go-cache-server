// Package cache provides ...
package cache

import (
	"sync"
)

type inMemoryCache struct {
	c 		map[string][]byte
	mutex	sync.RWMutex
	Stat
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, exist := c.c[k]
	// 更新 Stat 状态
	if exist {
		c.del(k, v)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		// 更新Stat状态
		c.del(k, v)
	}
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.c[k], nil
}

func (c *inMemoryCache) GetState() Stat {
	return c.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		make(map[string][]byte),
		sync.RWMutex{},
		Stat{}}
}
