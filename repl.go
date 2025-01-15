package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tylerbartlett24/pokedex/internal/pokeapi"
	"github.com/tylerbartlett24/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
	cache            pokecache.Cache
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		parameter := ""
		if len(words) > 1 {
			parameter = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, parameter)
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
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name: "explore",
			description: "Takes area as argument, returns list of pokemon in" +
				"area.",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Takes pokemon name as argument, attempts to catch" +
				"it.",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Takes captured pokemon name as argument, prints " +
				"name, height, weight, stats and types.",
			callback: commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints list of pokemon the user has caught.",
			callback:    commandPokedex,
		},
	}
}
