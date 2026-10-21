package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)

type LocArea struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}


func ListLocAreas(cache *pokecache.Cache, pageURL string) (LocArea, error) {
	value, exists := cache.Get(pageURL)
	if exists {
		// The line below is just for debugging purposes
		fmt.Println("Using cached data for", pageURL)

	} else {
		// The line below is just for debugging purposes
		fmt.Println("Fetching fresh data for", pageURL)
		
		res, err := http.Get(pageURL)
		if err != nil {
			return LocArea{}, fmt.Errorf("failed to fetch location areas: %w", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocArea{}, fmt.Errorf("failed to read response body: %w", err)
		}
		if res.StatusCode > 299 {
			return LocArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		value = body
		cache.Add(pageURL, body)
	}

	areas, err := UnmarshalLocAreas(value)
	if err != nil {
		return LocArea{}, fmt.Errorf("failed to unmarshal location areas: %w", err)
	}
	return areas, nil
}


func UnmarshalLocAreas(data []byte) (LocArea, error) {
	var areas LocArea
	if err := json.Unmarshal(data, &areas); err != nil {
		return LocArea{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return areas, nil
}