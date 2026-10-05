package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PancakeMarquis/pokedexcli.git/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	prevLocation  *string
	nextLocation  *string
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println(err)
		}
		text := scanner.Text()

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
			continue
		} else {
			fmt.Printf("Unknown command\n")
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
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show prevouis locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore encounters in area",
			callback:    commandExplore,
		},
	}
	return commands
}
