package lib

import (
	"errors"
	"fmt"
	"sync"
)

type CacheContainer struct {
	CacheFactory     map[string]*[]byte
	PersistenceMode  bool
	LocalStoragePath string
	Mtx              sync.RWMutex
}

func NewCacheContainer(persistence bool, localpath string) *CacheContainer {
	return &CacheContainer{
		CacheFactory:     map[string]*[]byte{},
		PersistenceMode:  persistence,
		LocalStoragePath: localpath,
	}
}

func (c *CacheContainer) AddNewCacheKey(cacheKey string) error {
	if cacheKey == "" {
		return errors.New("cache key is required")
	}
	defer c.Mtx.Unlock()
	c.Mtx.Lock()

	if _, ok := c.CacheFactory[cacheKey]; ok {
		return errors.New("cache key already exist")
	}

	c.CacheFactory[cacheKey] = &[]byte{}
	return nil
}

func (c *CacheContainer) AppendDataToCache(cacheKey string, data []byte) error {
	if cacheKey == "" {
		return errors.New("cache key is required")
	}

	if len(data) == 0 {
		return errors.New("cache data cannot be empty")
	}

	defer c.Mtx.Unlock()
	c.Mtx.Lock()

	if _, ok := c.CacheFactory[cacheKey]; !ok {
		return fmt.Errorf("cache with key %s not found", cacheKey)
	}

	if c.CacheFactory == nil {
		c.CacheFactory = map[string]*[]byte{}
	}

	c.CacheFactory[cacheKey] = &data
	return nil
}

func (c *CacheContainer) GetCacheByKey(key string) (*[]byte, error) {
	if key == "" {
		return nil, errors.New("cache key is required")
	}
	defer c.Mtx.RLock()
	c.Mtx.RUnlock()

	data, ok := c.CacheFactory[key]
	if !ok {
		return nil, fmt.Errorf("cache with key %s not found", key)
	}

	return data, nil
}
