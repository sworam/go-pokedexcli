package main

import (
	"fmt"
	"github.com/sworam/go-pokedexcli/internal/pokeapi"
)

func commandInspect(c *config, a ...string) error {
	pokemonName := a[0]
	pokemon, ok := c.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("Pokemon with name '%s' not found in pokedex", pokemonName)
	}
	displayPokemon(pokemon)
	return nil
}

func displayPokemon(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  -%s\n", pokeType.Type.Name)
	}
}
