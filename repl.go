package main

import(
	"strings"
)

func cleanInput(text string) []string {
	splitText := strings.Fields(text)

	result := []string{}
	for _, word := range splitText {
		result = append(result, strings.ToLower(word))
	}
	
	return result
}