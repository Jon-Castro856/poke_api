package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetData(cmd string) ([]byte, error) {
	if cmd == "map" {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area")
		if err != nil {
			return nil, err
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

		return body, nil
	}
	return nil, nil
}

func ProcessData(data []byte) error {
	pokeMap := MapData{}
	if err := json.Unmarshal(data, &pokeMap); err != nil {
		return err
	}
	fmt.Println(pokeMap.Results[0].Name)
	return nil
}

type MapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
