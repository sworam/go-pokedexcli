package main

import (
	"fmt"
)

func commandPokedex(c *config, a ...string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range c.pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
