package main

import (
	"fmt"
	"strings"
)

func ehPalindromo(str string) bool {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if ehPalindromo(input) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}
