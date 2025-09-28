package main

import (
	"bufio"
	"fmt"
	"github.com/sworam/go-pokedexcli/internal/pokeapi"
	"github.com/sworam/go-pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

type config struct {
	next     string
	previous string
	cache    pokecache.Cache
	pokedex  []pokeapi.Pokemon
}

var commandRegistry = map[string]cliCommand{}

func main() {
	registerCommands()
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		cache: pokecache.NewCache(time.Second * 5),
	}
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
			c.callback(&conf, cleaned[1:]...)
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
