package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("frase 1: ")
	frase1 := bufio.NewScanner(os.Stdin)
	frase1.Scan()
	fmt.Println("frase 2: ")
	frase2 := bufio.NewScanner(os.Stdin)
	frase2.Scan()

	if strings.Contains(frase1.Text(), frase2.Text()) {
		fmt.Println("substring")
	} else {
		fmt.Println("nao")
	}

}
