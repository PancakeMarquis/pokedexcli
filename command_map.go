package main

import (
	"fmt"

	"github.com/PancakeMarquis/pokedexcli.git/internal/pokeapi"
)

func commandMap(cfg *config) error {
	var url *string
	if cfg.nextLocation != nil {
		url = cfg.nextLocation
	}
	client := pokeapi.Client{}

	locations, err := client.ListLocations(url)
	if err != nil {
		return err
	}

	for _, item := range locations.Results {
		fmt.Println(item.Name)
	}
	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous

	return nil
}
