// Contagem de vogais:
// Escreva uma função que receba uma string como entrada e retorne o número de vogais na string.
package main

import (
	"fmt"
	"strings"
)

func contarVogais(texto string) int {
	vogais := "aeiouAEIOU"
	contador := 0

	for _, char := range texto {
		if strings.ContainsAny(string(char), vogais) {
			contador++
		}
	}

	return contador
}

func main() {
	resultado := contarVogais("Hello World")
	fmt.Printf("O número de vogais no texto é %d", resultado)
}
