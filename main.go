package main

import (
	"time"

	"github.com/PancakeMarquis/pokedexcli.git/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.PokemonJson),
	}
	startRepl(&cfg)
}
