package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex! \n")
	fmt.Printf("Usage: \n")
	fmt.Printf("\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Printf("\n")
	return nil
}
