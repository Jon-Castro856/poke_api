package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jon-Castro856/poke_api/internal/api"
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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Explains how to use the pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Provides a list of 20 in game areas, subsequent calls provide the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapBack",
			description: "Provides previous list of in game areas",
			callback:    commandMapBack,
		},
		"test": {
			name:        "test",
			description: "test",
			callback:    commandTest,
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
	return nil
}

func commandMapBack() error {
	return nil
}

func commandTest() error {
	api.GetData()
	return nil
}
