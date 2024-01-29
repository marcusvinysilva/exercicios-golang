// Soma de dois números:
// Escreva uma função que receba dois números como entrada e retorne a soma deles.
package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func main() {
	resultado := soma (5, 7)
	fmt.Println(resultado)
}
