package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocArea struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

//var BaseURL *string = ("https://pokeapi.co/api/v2/location-area/")

func ListLocAreas(pageURL string) (LocArea, error) {
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

	var areas LocArea
	if err := json.Unmarshal(body, &areas); err != nil {
		return LocArea{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return areas, nil
}