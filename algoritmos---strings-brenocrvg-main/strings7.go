package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func containsNumber(input string) bool {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma string: ")
	input, _ := reader.ReadString('\n')

	if containsNumber(input) {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}
