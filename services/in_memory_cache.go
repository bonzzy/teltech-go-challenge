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
	Cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))

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
