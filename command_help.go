package main

import "fmt"

func commandHelp(cfg *config, parameters ...string) error {
	fmt.Printf("Welcome to the Pokedex! \n")
	fmt.Printf("Usage: \n")
	fmt.Printf("\n")
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Printf("\n")
	return nil
}
