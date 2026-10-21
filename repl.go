package main

import (
	"fmt"
	"os"
	"strings"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"time"
)

var cache = pokecache.NewCache(10 * time.Second)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
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

func commandMapb(cfg *config) error {
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