package main

import (
	"fmt"
	"strings"
)

func main() {
	var legal string
	fmt.Print("digite: ")
	fmt.Scan(&legal)

	mexer := strings.ToUpper(legal)
	fmt.Println("eh isso: ", mexer)

}
