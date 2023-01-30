package services

import (
	"github.com/allegro/bigcache"
	"time"
)

var Cache *bigcache.BigCache

func GetCache() *bigcache.BigCache {
	if Cache != nil {
		return Cache
	}

	CacheConfig := bigcache.DefaultConfig(time.Minute)
	// Check what needs to be evicted every 1Second
	// meaning that it might pass max minute and 1 second to evict the value instead of exactly 1 min
	// BigCache advises not to set < 1 Second
	CacheConfig.CleanWindow = time.Second
	Cache, _ = bigcache.NewBigCache(CacheConfig)

	return Cache
}

func SetCacheValue(key string, value string) {
	err := GetCache().Set(key, []byte(value))
	if err != nil {
		return
	}
}

func GetCacheValue(key string) (string, error) {
	rawValue, err := GetCache().Get(key)

	return string(rawValue), err
}
