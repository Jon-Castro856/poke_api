package main

import (
	"fmt"
	"os"

	"github.com/Jon-Castro856/poke_api/internal/api"
	"github.com/Jon-Castro856/poke_api/internal/structs"
)

func getCommands() map[string]structs.CliCommand {
	return map[string]structs.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Explains how to use the pokedex",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Provides a list of 20 in-game areas, subsequent calls provide the next 20 areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "map back",
			Description: "Provides previous list of in-game areas",
			Callback:    commandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "explore a given location",
			Callback:    commandExplore,
		},
	}
}

func commandExit(cfg *structs.Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

func commandHelp(cfg *structs.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println()
	fmt.Printf("config is currently %s\n and %s\n", cfg.Forward, cfg.Back)
	return nil
}

func commandMap(cfg *structs.Config) error {
	mapInfo, err := api.GetData(cfg.Forward, cache)
	if err != nil {
		fmt.Printf("error acquiring data")
	}

	mapList, err := api.ProcessData(mapInfo)
	if err != nil {
		fmt.Print("error processing data")
	}

	cfg.Back = mapList.Previous
	cfg.Forward = mapList.Next

	for _, area := range mapList.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapBack(cfg *structs.Config) error {
	mapInfo, err := api.GetData(cfg.Back, cache)
	if err != nil {
		fmt.Println("error acquiring data")
	}
	mapList, err := api.ProcessData(mapInfo)
	if err != nil {
		fmt.Println("error processing data")
	}

	cfg.Back = mapList.Previous
	cfg.Forward = mapList.Next

	for _, area := range mapList.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cfg *structs.Config)) error {
	return nil
}
