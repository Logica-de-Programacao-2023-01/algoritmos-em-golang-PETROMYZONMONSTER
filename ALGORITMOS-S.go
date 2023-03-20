//Faça um algoritmo que leia dois números inteiros e mostre o maior deles.

package main

import "fmt"

func main() {
	var n1 float64
	var n2 float64
	fmt.Println("Digite o seu primeiro número:")
	fmt.Scan(&n1)
	fmt.Println("Digite o seu segundo número:")
	fmt.Scan(&n2)
	if n1 > n2 {
		fmt.Println("O maior entre esses números é:", n1)
	} else {
		fmt.Println("O maior entre esses números é:", n2)
	}
}
