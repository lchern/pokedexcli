package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		for _, command := range getCommands() {
			if firstWord == command.name {
				isCommandFound = true
				err := command.callback()
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
