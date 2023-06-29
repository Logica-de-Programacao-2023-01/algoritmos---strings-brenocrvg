package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagrama(str1ng, strong string) bool {
	str1ng = strings.ReplaceAll(str1ng, " ", "")
	strong = strings.ReplaceAll(strong, " ", "")
	str1ng = strings.ToLower(str1ng)
	strong = strings.ToLower(strong)

	if len(str1ng) != len(strong) {
		return false
	}
	chars1 := strings.Split(str1ng, "")
	chars2 := strings.Split(strong, "")
	sort.Strings(chars1)
	sort.Strings(chars2)
	for i := 0; i < len(chars1); i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}

	return true
}

func main() {
	var input1, input2 string
	fmt.Print("digite a primeira string: ")
	fmt.Scanln(&input1)
	fmt.Print("digite a segunda string: ")
	fmt.Scanln(&input2)

	if anagrama(input1, input2) {
		fmt.Println("as strings são anagramas.")
	} else {
		fmt.Println("as strings não são anagramas.")
	}
}
