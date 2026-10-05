package main

import (
	"fmt"
)

func commandMap(cfg *config, parameters ...string) error {
	var url *string
	if cfg.nextLocation != nil {
		url = cfg.nextLocation
	}

	locations, err := cfg.pokeapiClient.ListLocations(url)
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
