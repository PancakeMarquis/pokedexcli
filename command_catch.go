package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("Please add a name or id")

	}
	name := parameters[0]

	pokemon, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}
	challengeNum := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if challengeNum > 40 {
		fmt.Printf("%s escaped\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
