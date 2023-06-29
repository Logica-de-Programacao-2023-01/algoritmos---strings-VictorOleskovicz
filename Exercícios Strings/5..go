package main

import (
	"fmt"
	"strconv"
)

func verificarNumeroFlutuante(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if verificarNumeroFlutuante(input) {
		fmt.Println("É um número válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um número válido em ponto flutuante.")
	}
}
