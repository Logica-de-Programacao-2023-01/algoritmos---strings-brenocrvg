package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	legal := strings.Split(str, "")

	for i := 0; i < len(legal)/2; i++ {
		j := len(legal) - i - 1
		legal[i], legal[j] = legal[j], legal[i]
	}
	return strings.Join(legal, "")
}

func main() {
	var frase string
	fmt.Print("digite uma string: ")
	fmt.Scanln(&frase)

	inverter := reverseString(frase)
	fmt.Println("frase:", inverter)
}
