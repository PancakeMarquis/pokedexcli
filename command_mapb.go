package main

import (
	"fmt"

	"github.com/PancakeMarquis/pokedexcli.git/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocation == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := cfg.prevLocation

	client := pokeapi.Client{}
	locations, err := client.ListLocations(url)
	if err != nil {
		return err
	}
	for _, item := range locations.Results {
		fmt.Println(item.Name)
	}
	cfg.prevLocation = locations.Previous
	cfg.nextLocation = locations.Next

	return nil
}
