package main

import (
	"fmt"
	"github.com/sworam/go-pokedexcli/internal/pokeapi"
	"math/rand"
)

func commandCatch(c *config, a ...string) error {
	pokemonName := a[0]
	pokemon, err := pokeapi.GetPokemon(pokemonName, &c.cache)
	if err != nil {
		return err
	}
	tryToCatchPokemon(c, pokemon)
	return nil
}

func tryToCatchPokemon(c *config, pokemon pokeapi.Pokemon) {
	catchRate := calcCatchRate(pokemon)

	if rand.Float64() <= catchRate {
		fmt.Printf("Caught '%s'\n", pokemon.Name)
		c.pokedex = append(c.pokedex, pokemon)
	} else {
		fmt.Printf("'%s' escaped!\n", pokemon.Name)
	}
}

func calcCatchRate(pokemon pokeapi.Pokemon) float64 {
	catchRate := 1.9 / (float64(pokemon.BaseExperience) / 20)
	return catchRate
}
