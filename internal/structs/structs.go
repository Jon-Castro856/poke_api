package structs

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(Cfg *Config) error
}

type Config struct {
	ApiClient Client
	Back      string
	Forward   string
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

type Client struct {
	HttpClient http.Client
	Cache      Cache
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	data, ok := c.Data[key]

	return data.Val, ok
}

func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Data[key] = CacheEntry{
		Val:       val,
		CreatedAt: time.Now().UTC(),
	}
	fmt.Println("added to cache. total cache entries:")
	fmt.Println(len(c.Data))
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.Reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) Reap(now time.Time, last time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for key, entry := range c.Data {
		if entry.CreatedAt.Before(now.Add(-last)) {

			fmt.Println("clearing data...")
			delete(c.Data, key)
		}
	}

}
