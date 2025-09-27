package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetLocation(url string) (Location, error) {
	var finalURL = "https://pokeapi.co/api/v2/location-area"
	if url != "" {
		finalURL = url
	}

	res, err := http.Get(finalURL)
	if err != nil {
		return Location{}, fmt.Errorf("Connection Error")
	}

	defer res.Body.Close()
	var location Location
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&location); err != nil {
		return Location{}, fmt.Errorf("Could not decode json")
	}
	return location, nil
}
