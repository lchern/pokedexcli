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

type PokemonInArea struct {
	Results []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name   string `json:"name"`
	BaseXP int `json:"base_experience"`
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

func ListPokemonInArea(cache *pokecache.Cache, areaName string) (PokemonInArea, error) {
	areaURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)
	
	value, exists := cache.Get(areaURL)
	if exists {
		// The line below is just for debugging purposes
		fmt.Println("Using cached data for", areaURL)
		
	} else {
		// The line below is just for debugging purposes
		fmt.Println("Fetching fresh data for", areaURL)
		
		res, err := http.Get(areaURL)
		if err != nil {
			return PokemonInArea{}, fmt.Errorf("failed to fetch Pokemon in area %s: %w", areaName, err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return PokemonInArea{}, fmt.Errorf("failed to read response body: %w", err)
		}
		if res.StatusCode > 299 {
			return PokemonInArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		value = body
		cache.Add(areaURL, body)
	}

	pokemon, err := UnmarshalPokemonInArea(value)
	if err != nil {
		return PokemonInArea{}, fmt.Errorf("failed to unmarshal Pokemon in area %s: %w", areaName, err)
	}
	return pokemon, nil
}

func UnmarshalPokemonInArea(data []byte) (PokemonInArea, error) {
	var pokemon PokemonInArea
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return PokemonInArea{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return pokemon, nil
}

func GetPokemonDetails(cache *pokecache.Cache, pokemonName string) (Pokemon, error) {
	pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)

	value, exists := cache.Get(pokemonURL)
	if exists {
		// The line below is just for debugging purposes
		fmt.Println("Using cached data for", pokemonURL)

	} else {
		// The line below is just for debugging purposes
		fmt.Println("Fetching fresh data for", pokemonURL)

		res, err := http.Get(pokemonURL)
		if err != nil {
			return Pokemon{}, fmt.Errorf("failed to fetch details for Pokemon %s: %w", pokemonName, err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("failed to read response body: %w", err)
		}
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		value = body
		cache.Add(pokemonURL, body)
	}

	pokemon, err := UnmarshalPokemonData(value)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to unmarshal Pokemon details for %s: %w", pokemonName, err)
	}
	return pokemon, nil
}

func UnmarshalPokemonData(data []byte) (Pokemon, error) {
	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return pokemon, nil
}