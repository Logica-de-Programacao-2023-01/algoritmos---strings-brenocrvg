package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := substituit(input)
	fmt.Println("Resultado:", result)
}

func substituit(input string) string {
	vogais := "aeiouAEIOU"
	result := strings.Builder{}

	for _, top := range input {
		if strings.ContainsRune(vogais, top) {
			result.WriteString("*")
		} else {
			result.WriteRune(top)
		}
	}

	return result.String()
}
