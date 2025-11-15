package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Jon-Castro856/poke_api/internal/pokecache"
	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeClient := structs.Client{
		HttpClient: http.Client{
			Timeout: 5 * time.Second,
		},
		Cache: pokecache.NewCache(1 * time.Minute),
	}
	cfg := &structs.Config{
		ApiClient: pokeClient,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanUp := cleanInput((line))
		if len(cleanUp) == 0 {
			continue
		}
		commandName := cleanUp[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
