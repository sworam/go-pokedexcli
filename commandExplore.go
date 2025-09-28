package main

import (
	"fmt"
	"github.com/sworam/go-pokedexcli/internal/pokeapi"
)

func commandExplore(c *config, a ...string) error {
	locName := a[0]
	detailedLocation, err := pokeapi.GetDetailedLocation(locName, &c.cache)
	if err != nil {
		return err
	}
	showPokemonOfDetailedLocation(detailedLocation)
	return nil
}

func showPokemonOfDetailedLocation(detLocation pokeapi.DetailedLocation) {
	for _, encounter := range detLocation.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}
