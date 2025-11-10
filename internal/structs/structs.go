package structs

import (
	"sync"
	"time"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(Cfg *Config, cache *Cache) error
}

type Config struct {
	Back    string
	Forward string
}

type MapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Data     map[string]CacheEntry
	Mutex    *sync.Mutex
	Interval time.Duration
}

func (c Cache) Get(key string) ([]byte, bool) {
	data := c.Data[key].Val
	if data != nil {
		return data, true
	} else {
		return nil, false
	}
}

func (c Cache) Add(key string, val []byte) {
	newEntry := CacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
	c.Data[key] = newEntry
}

func (c Cache) ReapLoop() {
	ticker := time.NewTicker(c.Interval)

	for range ticker.C {
		for key, entry := range c.Data {
			if time.Since(entry.CreatedAt) > c.Interval {
				delete(c.Data, key)
			}
		}
	}
}
