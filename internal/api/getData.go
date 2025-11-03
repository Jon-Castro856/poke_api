package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func GetData(cmd string, cfg structs.Config) ([]byte, error) {
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
