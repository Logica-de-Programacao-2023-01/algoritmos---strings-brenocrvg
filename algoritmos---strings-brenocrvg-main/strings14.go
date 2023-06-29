package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if decrescente(input) {
		fmt.Println("decrescente")
	} else {
		fmt.Println("não decrescente")
	}
}

func decrescente(input string) bool {
	if len(input) < 2 {
		return false
	}

	for i := 0; i < len(input)-1; i++ {
		curr, _ := strconv.Atoi(string(input[i]))
		next, _ := strconv.Atoi(string(input[i+1]))

		if curr <= next {
			return false
		}
	}

	return true
}
