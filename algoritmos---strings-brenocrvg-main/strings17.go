package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("D: ")
	fmt.Scanln(&input)

	letrasUnicas := encontrar(input)
	fmt.Println("únicas:", letrasUnicas)
}
func encontrar(input string) string {
	frequencia := make(map[rune]int)
	for _, char := range input {
		frequencia[char]++
	}
	letrasUnicas := ""
	for _, char := range input {
		if frequencia[char] == 1 {
			letrasUnicas += string(char)
		}
	}
	return letrasUnicas
}
