package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Combtz/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient *pokeapi.Client
	next          string
	previous      string
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "show 20 maps at a time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "goes back a page show 20 maps at a time",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Founds pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
}
