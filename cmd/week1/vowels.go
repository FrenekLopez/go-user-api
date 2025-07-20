package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	s = strings.ToLower(s)
	count := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}
	return count
}

func main() {
	text := "Hola Mundo"
	result := countVowels(text)
	fmt.Printf("Número de vocales en '%s': %d\n", text, result)
}
