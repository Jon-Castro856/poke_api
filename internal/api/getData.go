package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func GetData(cmd string, offsetUrl string) ([]byte, error) {
	fmt.Printf("the command is %s", cmd)
	var url string
	if cmd == "map" {
		if offsetUrl != "" {
			url = offsetUrl
		} else {
			url = "https://pokeapi.co/api/v2/location-area"
		}
	}

	if cmd == "map back" {
		if offsetUrl != "" {
			url = offsetUrl
		} else {
			url = "https://pokeapi.co/api/v2/location-area"
		}
	}
	fmt.Printf("The Current url is %s\n", url)
	if url == "" {
		fmt.Println("URL value empty, list is either at the start or end")
	}

	res, err := http.Get(url)
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

func ProcessData(data []byte) (structs.MapData, error) {
	pokeMap := structs.MapData{}
	if err := json.Unmarshal(data, &pokeMap); err != nil {
		return structs.MapData{}, err
	}
	return pokeMap, nil
}
