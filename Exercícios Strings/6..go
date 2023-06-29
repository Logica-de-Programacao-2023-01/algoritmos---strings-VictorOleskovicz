package main

import (
	"fmt"
	"strings"
)

func contarPalavras(str string) int {
	palavras := strings.Fields(str)
	return len(palavras)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	numPalavras := contarPalavras(input)
	fmt.Println("A string contém", numPalavras, "palavra(s).")
}
