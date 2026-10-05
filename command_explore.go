package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("Please add a name or id")

	}
	name := parameters[0]

	encounters, err := cfg.pokeapiClient.ListEncounters(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon: ")
	for _, item := range encounters.PokemonEncounters {
		fmt.Println(item.Pokemon.Name)
	}

	return nil
}
