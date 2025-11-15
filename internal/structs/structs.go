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
	Command   []string
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

type LocationDetail struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// Methods----------------------------------------------
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
