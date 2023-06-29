package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(input string) int {
	words := strings.Fields(input)
	return len(words)
}

func main() {
	frase := bufio.NewScanner(os.Stdin)
	fmt.Print("frase: ")
	frase.Scan()
	input := frase.Text()
	numWords := countWords(input)
	fmt.Println("Número de palavras:", numWords)
}

