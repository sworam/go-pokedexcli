package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sworam/go-pokedexcli/internal/pokecache"
)

func GetPokemon(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {
	finalURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	var pokemon Pokemon

	err := getFromCacheOrRequest(finalURL, cache, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}
	return pokemon, nil
}

func GetDetailedLocation(locName string, cache *pokecache.Cache) (DetailedLocation, error) {
	finalURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locName)
	var location DetailedLocation

	err := getFromCacheOrRequest(finalURL, cache, &location)
	if err != nil {
		return DetailedLocation{}, err
	}
	return location, nil
}

func GetLocation(url string, cache *pokecache.Cache) (Location, error) {
	var finalURL = "https://pokeapi.co/api/v2/location-area"
	if url != "" {
		finalURL = url
	}

	var location = Location{}
	err := getFromCacheOrRequest(finalURL, cache, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}

func getFromCacheOrRequest(url string, cache *pokecache.Cache, v any) error {
	data, ok := cache.Get(url)
	if !ok {
		data, err := makeGETRequest(url)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	data, _ = cache.Get(url)
	err := json.Unmarshal(data, &v)
	return err
}

func makeGETRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Connection Error")
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
