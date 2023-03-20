//Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.

package main

import "fmt"

func main() {
	var n1 int
	fmt.Println("Digite o seu número:")
	fmt.Scan(&n1)
	if n1%2 == 0 {
		fmt.Println("Seu número é par.")
	} else {
		fmt.Println("Seu número é impar")
	}
}
