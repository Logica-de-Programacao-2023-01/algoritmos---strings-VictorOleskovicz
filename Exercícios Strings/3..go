package main

import (
	"fmt"
	"strings"
)

func substituirCaractere(str string, caractereAntigo string, caractereNovo string) string {
	return strings.ReplaceAll(str, caractereAntigo, caractereNovo)
}

func main() {
	var input string
	var caractereAntigo string
	var caractereNovo string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&caractereAntigo)
	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&caractereNovo)

	resultado := substituirCaractere(input, caractereAntigo, caractereNovo)
	fmt.Println("Resultado:", resultado)
}
