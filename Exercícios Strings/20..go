package main

import (
	"fmt"
	"regexp"
	"strings"
)

func verificaCamelCase(str string) (bool, int) {
	regex := regexp.MustCompile("^[a-z]+(?:[A-Z][a-z]+)*$")
	words := strings.FieldsFunc(str, func(r rune) bool {
		return r == '_' || r == '-'
	})
	return regex.MatchString(str), len(words)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	isCamelCase, numPalavras := verificaCamelCase(input)

	if isCamelCase {
		fmt.Println("A string está em camelCase.")
		fmt.Println("Número de palavras:", numPalavras)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
