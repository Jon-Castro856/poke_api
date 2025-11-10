package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func GetData(offsetUrl string, cache *structs.Cache) ([]byte, error) {
	var url string
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if offsetUrl != "" {
		url = offsetUrl
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	if url == "" {
		fmt.Println("URL value empty, list is either at the start or end")
	}

	data, ok := cache.Get(url)
	if ok {
		fmt.Println("data found in the cache")
		return data, nil
	}
	fmt.Println("no data in the cache, making api call")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error with completing request: %v", err)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		fmt.Printf("Erorr retreiving data: %v", res.StatusCode)
	}
	fmt.Println("response succesful")

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	cache.Add(url, body)

	return body, nil
}

func ProcessData(data []byte) (structs.MapData, error) {
	pokeMap := structs.MapData{}
	if err := json.Unmarshal(data, &pokeMap); err != nil {
		return structs.MapData{}, err
	}
	return pokeMap, nil
}
