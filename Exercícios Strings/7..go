package main

import (
	"fmt"
	"regexp"
)

func contemNumero(str string) bool {
	regex := regexp.MustCompile("[0-9]")
	return regex.MatchString(str)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if contemNumero(input) {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}
