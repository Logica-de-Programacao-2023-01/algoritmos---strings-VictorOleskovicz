package main

import (
	"fmt"
	"strings"
)

func substituirLetra(str string, letraAntiga string, letraNova string) string {
	return strings.ReplaceAll(str, letraAntiga, letraNova)
}

func main() {
	var input string
	var letraAntiga string
	var letraNova string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)
	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&letraAntiga)
	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&letraNova)

	resultado := substituirLetra(input, letraAntiga, letraNova)
	fmt.Println("Resultado:", resultado)
}
