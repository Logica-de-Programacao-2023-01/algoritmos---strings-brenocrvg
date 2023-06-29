package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	frase1 := bufio.NewScanner(os.Stdin)
	frase2 := bufio.NewScanner(os.Stdin)
	fmt.Print("digite a frase 1: ")
	frase1.Scan()
	fmt.Print("digite a frase 2: ")
	frase2.Scan()

	if frase1 == frase2 {
		fmt.Println("as frases sao iguais")
	} else {
		fmt.Println("as frases sao diferenmtes")
	}

}
