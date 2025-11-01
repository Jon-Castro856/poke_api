package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			//page:        &config{},
		},
		"help": {
			name:        "help",
			description: "Explains how to use the pokedex",
			callback:    commandHelp,
			//page:        &config{},
		},
		"map": {
			name:        "map",
			description: "Provides a list of 20 in-game areas, subsequent calls provide the next 20 areas",
			callback:    commandMap,
			//page:        &config{},
		},
		"mapb": {
			name:        "map back",
			description: "Provides previous list of in-game areas",
			callback:    commandMapBack,
			//page:        &config{},
		},
		"test": {
			name:        "test",
			description: "test",
			callback:    commandTest,
			//page:        &config{},
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	urlBack    string
	urlForward string
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
