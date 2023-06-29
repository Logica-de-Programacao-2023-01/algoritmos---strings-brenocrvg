package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("digite uma string: ")
	fmt.Scanln(&input)

	resultado := remover(input)
	fmt.Println("resultado:", resultado)
}

func remover(input string) string {
	vogais := "aeiouAEIOU"
	result := strings.Builder{}

	for _, massa := range input {
		if !strings.ContainsRune(vogais, massa) {
			result.WriteRune(massa)
		}
	}

	return result.String()
}
