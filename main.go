package main

import (
	"bufio"
	"fmt"
	"github.com/sworam/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	next     string
	previous string
}

var commandRegistry = map[string]cliCommand{}

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
}

func main() {
	registerCommands()
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	var conf config
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		command := cleaned[0]
		c, ok := commandRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			c.callback(&conf)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var cleanedWords []string
	for _, word := range words {
		trimmed := strings.Trim(word, " ")
		if len(trimmed) > 0 {
			cleanedWords = append(cleanedWords, strings.ToLower(trimmed))
		}
	}
	return cleanedWords
}

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
