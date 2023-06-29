package main

import (
	"fmt"
	"regexp"
)

func contemApenasNumeros(str string) bool {
	regex := regexp.MustCompile("^[0-9]+$")
	return regex.MatchString(str)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if contemApenasNumeros(input) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string contém caracteres não numéricos.")
	}
}
