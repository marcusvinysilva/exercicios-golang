// Maior entre três números:Escreva uma função que receba três números como entrada e retorne o maior deles.
package main

import "fmt"

func maiorEntreTres(a, b, c int) interface{} {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else if c > a && c > b {
		return c
	} else {
		return "Nenhum número é maior"
	}
}

func main() {
	resultado := maiorEntreTres(8, 2, 5)
	fmt.Println(resultado)
}
