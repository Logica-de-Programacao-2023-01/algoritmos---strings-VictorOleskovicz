package main

import (
	"fmt"
	"strings"
)

func removerEspacos(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := removerEspacos(input)
	fmt.Println("Resultado:", resultado)
}
