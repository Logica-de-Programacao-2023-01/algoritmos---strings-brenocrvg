package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	isCamelCase := true
	wordCount := 1

	if len(input) > 0 && !unicode.IsUpper(rune(input[0])) {
		isCamelCase = false
	}

	for i := 1; i < len(input); i++ {
		if unicode.IsUpper(rune(input[i])) {
			if !unicode.IsLower(rune(input[i-1])) {
				isCamelCase = false
				break
			}
			wordCount++
		}
	}

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavras.\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
