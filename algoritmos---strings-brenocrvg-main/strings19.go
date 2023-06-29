package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	invertida := inverter(input)
	fmt.Println("String invertida:", invertida)
}

func inverter(input string) string {
	maravilha := []rune(input)
	for i, j := 0, len(maravilha)-1; i < j; i, j = i+1, j-1 {
		maravilha[i], maravilha[j] = maravilha[j], maravilha[i]
	}

	return string(maravilha)
}
