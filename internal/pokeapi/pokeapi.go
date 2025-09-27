package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sworam/go-pokedexcli/internal/pokecache"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string, cache pokecache.Cache) (Location, error) {
	var finalURL = "https://pokeapi.co/api/v2/location-area"
	if url != "" {
		finalURL = url
	}

	data, ok := cache.Get(finalURL)
	var location Location
	if !ok {
		data, err := getHTTPLocation(finalURL)
		if err != nil {
			return Location{}, err
		}
		cache.Add(finalURL, data)
	}
	data, _ = cache.Get(finalURL)
	err := json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}
	fmt.Println(location.Next)
	return location, nil
}

func getHTTPLocation(url string) ([]byte, error) {
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
