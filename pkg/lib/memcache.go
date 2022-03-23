package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type CacheContainer struct {
	CacheFactory     *map[string][]byte
	PersistenceMode  bool
	LocalStoragePath string
	Mtx              sync.RWMutex
}

func NewCacheContainer(persistence bool, localpath string) *CacheContainer {
	return &CacheContainer{
		CacheFactory:     &map[string][]byte{},
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

	if _, ok := (*c.CacheFactory)[cacheKey]; ok {
		return errors.New("cache key already exist")
	}

	(*c.CacheFactory)[cacheKey] = []byte{}
	return nil
}

func (c *CacheContainer) AddStructByKey(cacheKey string, data interface{}) error {
	if cacheKey == "" {
		return errors.New("cache key is required")
	}

	if data == nil {
		return errors.New("cache data cannot be empty")
	}

	//defer func() {
	//	log.Printf("\n\ncache data : %#v\n\n\n", c.CacheFactory)
	//}()
	defer c.Mtx.Unlock()
	c.Mtx.Lock()

	byteData, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	if c.CacheFactory == nil {
		c.CacheFactory = &map[string][]byte{}
	}

	(*c.CacheFactory)[cacheKey] = byteData
	return nil
}

func (c *CacheContainer) UpsertCacheByKey(cacheKey string, data []byte) error {
	if cacheKey == "" {
		return errors.New("cache key is required")
	}

	if len(data) == 0 {
		return errors.New("cache data cannot be empty")
	}

	defer c.Mtx.Unlock()
	c.Mtx.Lock()

	if _, ok := (*c.CacheFactory)[cacheKey]; !ok {
		return fmt.Errorf("cache with key %s not found", cacheKey)
	}

	if c.CacheFactory == nil {
		c.CacheFactory = &map[string][]byte{}
	}

	(*c.CacheFactory)[cacheKey] = data
	return nil
}

func (c *CacheContainer) GetStructByKey(key string, data interface{}) error {
	if key == "" {
		return errors.New("cache key is required")
	}

	if data == nil {
		return errors.New("struct can't be nil")
	}

	defer c.Mtx.RUnlock()
	c.Mtx.RLock()

	dataByte, ok := (*c.CacheFactory)[key]
	if !ok {
		return fmt.Errorf("cache with key %s not found", key)
	}

	err := json.Unmarshal(dataByte, data)
	return err
}

func (c *CacheContainer) GetCacheByKey(key string) (*[]byte, error) {
	if key == "" {
		return nil, errors.New("cache key is required")
	}
	defer c.Mtx.RUnlock()
	c.Mtx.RLock()

	data, ok := (*c.CacheFactory)[key]
	if !ok {
		return nil, fmt.Errorf("cache with key %s not found", key)
	}

	return &data, nil
}

func (c *CacheContainer) DeleteCacheByKey(key string) error {
	if key == "" {
		return errors.New("cache key is required")
	}
	defer c.Mtx.Unlock()
	c.Mtx.Lock()

	if _, ok := (*c.CacheFactory)[key]; !ok {
		return fmt.Errorf("cache with key %s not found", key)
	}

	delete(*c.CacheFactory, key)
	return nil
}
