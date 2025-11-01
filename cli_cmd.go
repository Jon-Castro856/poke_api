package main

import (
	"fmt"
	"os"

	"github.com/Jon-Castro856/poke_api/internal/api"
)

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap() error {
	mapInfo, err := api.GetData("map")
	if err != nil {
		fmt.Printf("error acquiring data")
	}
	api.ProcessData((mapInfo))
	return nil
}

func commandMapBack() error {
	mapInfo, err := api.GetData("mapb")
	if err != nil {
		fmt.Printf("error acquiring data")
	}
	api.ProcessData((mapInfo))
	return nil
}

func commandTest() error {
	api.GetData("map")
	return nil
}
