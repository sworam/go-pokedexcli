package main

import (
	"fmt"
	"github.com/sworam/go-pokedexcli/internal/pokeapi"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Usage:\n\n")

	for _, command := range commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(c *config) error {
	location, err := pokeapi.GetLocation(c.next)
	if err != nil {
		return err
	}
	displayLocation(location)
	c.next = location.Next
	c.previous = location.Previous
	return nil
}

func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	location, err := pokeapi.GetLocation(c.previous)
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
