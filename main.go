package main

import (
	"bufio"
	"fmt"
	"os"
)

var startPage = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

type config struct {
	commands map[string]cliCommand
	nextURL  *string
	prevURL  *string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var cfg = &config{
		commands: getCommands(),
		nextURL:  &startPage,
		prevURL:  nil,
	}	

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		firstWord := words[0]
		isCommandFound := false

		for _, command := range cfg.commands {
			if firstWord == command.name {
				isCommandFound = true
				err := command.callback(cfg)
				if err != nil {
					fmt.Println("Error executing command:", err)
				}
				break
			}

		}

		if !isCommandFound {
			fmt.Println("Unknown command")
		}
	}
}
