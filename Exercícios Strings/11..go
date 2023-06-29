package main

import (
	"fmt"
	"strings"
)

func removerVogais(str string) string {
	vogais := "aeiouAEIOU"
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(vogais, r) {
			return -1
		}
		return r
	}, str)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := removerVogais(input)
	fmt.Println("Resultado:", resultado)
}
