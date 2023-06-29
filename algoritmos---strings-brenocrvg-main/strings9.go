package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	frase := bufio.NewScanner(os.Stdin)
	fmt.Println("digite sua sacanagem: ")
	frase.Scan()

	var velha string
	var nova string
	fmt.Println("caractere para subristuir: ")
	fmt.Scan(&velha)
	fmt.Println("caractere novo no lugar: ")
	fmt.Scan(&nova)

	substituição := strings.Replace(frase.Text(), velha, nova, -1)
	fmt.Print(substituição)

}
