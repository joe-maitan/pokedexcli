package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"

	"github.com/joe-maitan/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient	 pokeapi.Client
	pokedex 	 	 map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
} // End config struct{}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
} // End cliCommand struct{}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
} // End cleanInput() func

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <name-of-area>",
			description: "Explore a area you discovered",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <name-of-pokemon>",
			description: "Attempt to catch a pokemon in the area you explored",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <name-of-pokemon>",
			description: "Inspect details of a pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print out all the pokemon you have caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
} // End getCommands() func

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
			fmt.Println("Unknown command")
			continue
		}
	}
} // End startRepl() func