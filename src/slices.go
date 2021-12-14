package main

import (
	"fmt"
	"strings"
)

func main() {

	slice := []string{"ola", "ke", "ase"}
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	fmt.Println(esPalindromo("ama"))
	//esto es falso pq compara strings por codigo asci y A no es igual que a
	fmt.Println(esPalindromo("Ama"))
	fmt.Println(esPalindromo("abale arroz a la zorra el abad"))
}

func esPalindromo(text string) bool {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	return strings.ToLower(text) == strings.ToLower(textReverse)
}
