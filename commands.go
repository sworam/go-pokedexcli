package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, a ...string) error
}

func registerCommands() {
	commandRegistry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commandRegistry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commandRegistry["map"] = cliCommand{
		name:        "map",
		description: "Displays the current location",
		callback:    commandMap,
	}
	commandRegistry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display the previous location",
		callback:    commandMapb,
	}
	commandRegistry["explore"] = cliCommand{
		name:        "explore",
		description: "Display all pokemon at a given location",
		callback:    commandExplore,
	}
	commandRegistry["catch"] = cliCommand{
		name:        "catch",
		description: "Try to catch a pokemon with the given name",
		callback:    commandCatch,
	}
	commandRegistry["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect a pokemon in the pokedex",
		callback:    commandInspect,
	}
}

func commandExit(c *config, a ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, a ...string) error {
	fmt.Print("Usage:\n\n")

	for _, command := range commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
