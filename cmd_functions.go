package main

import (
	"fmt"
	"os"

	"math/rand"

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

		"catch": {
			Name:        "catch",
			Description: "attempt to catch a pokemon and add it to your pokedex",
			Callback:    commandCatch,
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
	mapInfo, err := api.GetData(cfg.Forward, cfg)
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
	mapInfo, err := api.GetData(cfg.Back, cfg)
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

func commandExplore(cfg *structs.Config) error {
	if len(cfg.Command) != 2 {
		fmt.Println("enter the name of a location to explore")
		return nil
	}
	area := cfg.Command[1]

	areaUrl := "https://pokeapi.co/api/v2/location-area/" + area

	pokeInfo, err := api.GetData(areaUrl, cfg)
	if err != nil {
		fmt.Println("error acquiring data")
	}
	pokeList, err := api.ProcessLocData(pokeInfo)
	if err != nil {
		fmt.Println("error processing data")
	}
	for _, pokemon := range pokeList.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *structs.Config) error {
	if len(cfg.Command) != 2 {
		fmt.Println("enter the name of a pokemon to catch")
		return nil
	}
	pokemon := cfg.Command[1]

	pokeUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	monInfo, err := api.GetData(pokeUrl, cfg)
	if err != nil {
		fmt.Println("error acquiring data")
	}

	fmt.Printf("Attempting to catch %s...\n", pokemon)
	fmt.Printf("Throwing a pokeball at %s...\n", pokemon)

	monData, err := api.ProcessPokeData(monInfo)
	if err != nil {
		fmt.Println("error processing data")
		return nil
	}
	rate := float64((monData.BaseExperience / 600))
	catchRate := float64(int(rate*100)) / 100
	catchRate = min(0.9, max(catchRate, 0.25))
	result := rand.Float64() - catchRate

	fmt.Printf("base exp equals: %v\n", monData.BaseExperience)
	fmt.Printf("catch rate is %v, result is %v\n", catchRate, result)

	if result >= 1.0 {
		fmt.Printf("%s was succesfully caught!\n", pokemon)
		catch := structs.Pokemon{
			Name: pokemon,
			URL:  pokeUrl,
		}
		cfg.Caught[pokemon] = catch
	} else {
		fmt.Printf("failed to catch %s...\n", pokemon)
	}

	return nil
}
