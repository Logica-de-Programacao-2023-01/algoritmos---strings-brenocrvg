package main

import (
	"fmt"
	"strings"
)

func main() {
	//Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.
	frase := "manter fora do alcance de crianças"
	var velha string
	var nova string
	fmt.Println("caractere para subristuir: ")
	fmt.Scan(&velha)
	fmt.Println("caractere novo no lugar: ")
	fmt.Scan(&nova)

	substituição := strings.Replace(frase, velha, nova, -1)
	fmt.Print(substituição)

}
