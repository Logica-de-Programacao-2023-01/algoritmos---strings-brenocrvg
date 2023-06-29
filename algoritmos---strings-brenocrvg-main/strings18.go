package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Digite: ")
	fmt.Scanln(&input)

	if numeros(input) {
		fmt.Println("apenas números")
	} else {
		fmt.Println("não contém apenas números")
	}
}

func numeros(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
