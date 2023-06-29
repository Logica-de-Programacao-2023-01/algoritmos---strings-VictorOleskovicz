package main

import (
	"fmt"
	"strings"
)

func letrasUnicas(str string) string {
	uniques := ""
	for _, char := range str {
		if strings.Count(str, string(char)) == 1 {
			uniques += string(char)
		}
	}
	return uniques
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	uniques := letrasUnicas(input)
	fmt.Println("Letras únicas:", uniques)
}
