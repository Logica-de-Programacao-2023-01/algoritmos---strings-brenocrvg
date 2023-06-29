package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.
	frase := bufio.NewScanner(os.Stdin)
	fmt.Print("frase: ")
	frase.Scan()
	str := strings.ReplaceAll(frase.Text(), " ", "")
	fmt.Println(str)

}
