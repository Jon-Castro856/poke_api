package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func GetData(cmd string) error {
	url := ""

	if cmd == "test" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
