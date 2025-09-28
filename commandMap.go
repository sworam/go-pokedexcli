package main

import (
	"fmt"
	"github.com/sworam/go-pokedexcli/internal/pokeapi"
)

func commandMap(c *config, a ...string) error {
	location, err := pokeapi.GetLocation(c.next, &c.cache)
	if err != nil {
		return err
	}
	displayLocation(location)
	c.next = location.Next
	c.previous = location.Previous
	return nil
}

func commandMapb(c *config, a ...string) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	location, err := pokeapi.GetLocation(c.previous, &c.cache)
	if err != nil {
		return err
	}

	displayLocation(location)
	c.next = location.Next
	c.previous = location.Previous
	return nil
}

func displayLocation(location pokeapi.Location) {
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}
}
