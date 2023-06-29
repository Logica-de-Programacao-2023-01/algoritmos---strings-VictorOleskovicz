package main

import (
	"fmt"
	"strconv"
)

func ehSequenciaCrescente(str string) bool {
	for i := 1; i < len(str); i++ {
		prev, _ := strconv.Atoi(string(str[i-1]))
		curr, _ := strconv.Atoi(string(str[i]))
		if curr != prev+1 {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if ehSequenciaCrescente(input) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
