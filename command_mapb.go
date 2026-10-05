package main

import (
	"fmt"
)

func commandMapb(cfg *config, parameters ...string) error {
	if cfg.prevLocation == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := cfg.prevLocation

	locations, err := cfg.pokeapiClient.ListLocations(url)
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
