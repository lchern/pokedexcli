package main

import (
	"fmt"
	"os"
	"strings"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"time"
	"math/rand"
)

var cache = pokecache.NewCache(5 * time.Minute)

var pokedex = map[string]pokeapi.Pokemon{}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: 		 "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback: 	 commandMap,
		},
		"mapb": {
			name: 		 "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback: 	 commandMapb,
		},
		"explore": {
			name: 		 "explore",
			description: "Displays the names of all Pokemon found in that area",
			callback: 	 commandExplore,
		},
		"catch"	: {
			name: 		 "catch",
			description: "Catch a Pokemon in the current area",
			callback: 	 commandCatch,
		},
		"inspect": {
			name: 		 "inspect",
			description: "Inspect a Pokemon in your Pokedex",
			callback: 	 commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	splitText := strings.Fields(text)

	result := []string{}
	for _, word := range splitText {
		result = append(result, strings.ToLower(word))
	}

	return result
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {
	if cfg.nextURL == nil {
		fmt.Println("you're on the last page")
	} else {
		areas, err := pokeapi.ListLocAreas(cache, *cfg.nextURL)

		if err != nil {
			return fmt.Errorf("failed to fetch location areas: %w", err)
		}
		
		cfg.nextURL = areas.Next
		cfg.prevURL = areas.Previous
			
		for _, area := range areas.Results {
			fmt.Println(area.Name)	
		}
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
	} else {
		areas, err := pokeapi.ListLocAreas(cache, *cfg.prevURL)
	
		if err != nil {
			return fmt.Errorf("failed to fetch location areas: %w", err)
		}
		
		cfg.nextURL = areas.Next
		cfg.prevURL = areas.Previous
		
		for _, area := range areas.Results {
			fmt.Println(area.Name)	
		}
	}
	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name")
	}

	areaName := args[0]
	pokemon, err := pokeapi.ListPokemonInArea(cache, areaName)
	if err != nil {
		return fmt.Errorf("failed to fetch Pokemon in area %s: %w", areaName, err)
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, p := range pokemon.Results {
		fmt.Println("- " + p.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args []string) error {
	isCaught := false
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name to catch")
	}

	pokemonName := args[0]
	// Simulate catching the Pokemon
	pokemon, err := pokeapi.GetPokemonDetails(cache, pokemonName)
	if err != nil {
		return fmt.Errorf("failed to fetch Pokemon %s: %w", pokemonName, err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	time.Sleep(1 * time.Second) // Simulate some delay
	if rand.Intn(pokemon.BaseXP) < 30 { // Randomly determine if the Pokemon is caught based on its base experience
		isCaught = true
	}
	if isCaught {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}


func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name to inspect")
	}
	
	pokemonName := args[0]
	pokemon, exists := pokedex[pokemonName]
	if !exists {
		return fmt.Errorf("Pokemon %s not found in your Pokedex", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.Value)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}