package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("Please add a name or id")

	}
	name := parameters[0]

	pokemon, exist := cfg.pokedex[name]
	if !exist {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stats := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, types := range pokemon.Types {
		fmt.Printf("- %s\n", types.Type.Name)
	}

	return nil
}
