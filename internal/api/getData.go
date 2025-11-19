package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func GetData(offsetUrl string, config *structs.Config) ([]byte, error) {
	var url string

	if offsetUrl != "" {
		url = offsetUrl
	} else {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	data, ok := config.ApiClient.Cache.Get(url)
	if ok {
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := config.ApiClient.HttpClient.Do(req)
	if err != nil {
		fmt.Printf("error with completing request: %v", err)
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		fmt.Printf("Erorr retreiving data: %v", res.StatusCode)
		return nil, nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	config.ApiClient.Cache.Add(url, body)

	return body, nil
}

func ProcessData(data []byte) (structs.MapData, error) {
	pokeMap := structs.MapData{}
	if err := json.Unmarshal(data, &pokeMap); err != nil {
		return structs.MapData{}, err
	}
	return pokeMap, nil
}

func ProcessLocData(data []byte) (structs.LocationDetail, error) {
	pokeList := structs.LocationDetail{}
	if err := json.Unmarshal(data, &pokeList); err != nil {
		return structs.LocationDetail{}, err
	}
	return pokeList, nil
}

func ProcessPokeData(data []byte) (structs.PokemonData, error) {
	pokeData := structs.PokemonData{}
	if err := json.Unmarshal(data, &pokeData); err != nil {
		return structs.PokemonData{}, err
	}
	return pokeData, nil
}
