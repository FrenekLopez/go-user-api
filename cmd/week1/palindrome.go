package main

import (
	"fmt"
	"strings"
)

func Palindrome(s string) (bool, error) {
	if s == "" {
		return false, fmt.Errorf("Cadena vacia")
	}
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	word := "Holoh"
	result, err := Palindrome(word)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%s es palíndromo: %t\n", word, result)
}
