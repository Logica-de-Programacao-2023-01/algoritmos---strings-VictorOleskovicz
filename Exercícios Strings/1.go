package main

import (
	"fmt"
	"strings"
)

func converterParaMaiusculas(str string) string {
	return strings.ToUpper(str)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := converterParaMaiusculas(input)
	fmt.Println("Resultado:", resultado)
}
