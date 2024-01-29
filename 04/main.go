// Fatorial: Escreva uma função que calcule o fatorial de um número.
package main

import "fmt"

func fatorial(num uint) uint {
	if num == 0 {
		return 1
	} else {
		fat := uint(1)
		for i := num; i > 0; i-- {
			fat *= i
		}
		return fat
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
