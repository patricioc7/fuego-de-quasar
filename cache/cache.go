package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)
var mainCache *cache.Cache

func InitCache(){
	mainCache = cache.New(8*time.Minute, 10*time.Minute)
}

func SetValue(key string, value interface{}) {
	mainCache.Set(key, value,  cache.NoExpiration)
}

func GetValue(key string) (interface{}, bool) {
	return mainCache.Get(key)
}

