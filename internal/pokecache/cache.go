package pokecache

import (
	"sync"
	"time"

	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func NewCache(inter time.Duration) structs.Cache {
	cache := structs.Cache{
		Data:  map[string]structs.CacheEntry{},
		Mutex: &sync.Mutex{},
	}
	go cache.ReapLoop(inter)
	return cache
}
