package main

import (
	"fmt"
	"os"
	"bufio"
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
		fmt.Println("Your command was:", firstWord)
	}
}	