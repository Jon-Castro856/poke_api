package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Jon-Castro856/poke_api/internal/api"
	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanUp := cleanInput((line))
		if len(cleanUp) == 0 {
			continue
		}
		commandName := cleanUp[0]

		cfg := structs.Config{}

		command, exists := getCommands()[commandName]
		if exists {
			command.Cfg = cfg
			err := command.Callback(command.Name, command.Cfg)
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

func getCommands() map[string]structs.CliCommand {
	return map[string]structs.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
			Cfg:         structs.Config{},
		},
		"help": {
			Name:        "help",
			Description: "Explains how to use the pokedex",
			Callback:    commandHelp,
			Cfg:         structs.Config{},
		},
		"map": {
			Name:        "map",
			Description: "Provides a list of 20 in-game areas, subsequent calls provide the next 20 areas",
			Callback:    commandMap,
			Cfg:         structs.Config{},
		},
		"mapb": {
			Name:        "map back",
			Description: "Provides previous list of in-game areas",
			Callback:    commandMapBack,
			Cfg:         structs.Config{},
		},
	}
}

func commandExit(name string, cfg structs.Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

func commandHelp(name string, cfg structs.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println()
	return nil
}

func commandMap(name string, cfg structs.Config) error {
	mapInfo, err := api.GetData(name, cfg)
	if err != nil {
		fmt.Printf("error acquiring data")
	}
	processData((mapInfo))
	return nil
}

func commandMapBack(name string, cfg structs.Config) error {
	mapInfo, err := api.GetData(name, cfg)
	if err != nil {
		fmt.Printf("error acquiring data")
	}
	processData((mapInfo))
	return nil
}

func processData(data []byte) error {
	pokeMap := structs.MapData{}
	if err := json.Unmarshal(data, &pokeMap); err != nil {
		return err
	}

	for _, area := range pokeMap.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
