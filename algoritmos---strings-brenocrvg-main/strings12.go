package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if palindromo(input) {
		fmt.Println("palíndromo")
	} else {
		fmt.Println("nao")
	}
}

func palindromo(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}
