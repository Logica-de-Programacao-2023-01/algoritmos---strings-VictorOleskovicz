package main

import (
	"strings"
)

func substituirVogais(str string) string {
	vogais := "aeiouAEIOU"
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(vogais, r) {
			return '*'
		}
		return r
	}, str)
}

func main()
