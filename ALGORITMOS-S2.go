//Faça um algoritmo que leia três números inteiros e mostre o menor deles.

package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int
	fmt.Println("Digite o seu primeiro número:")
	fmt.Scan(&n1)
	fmt.Println("Digite o seu segundo número:")
	fmt.Scan(&n2)
	fmt.Println("Digite o seu terceiro número:")
	fmt.Scan(&n3)
	switch {
	case n1 < n2 && n1 < n3:
		fmt.Println("Seu menor número é:", n1)
	case n2 < n1 && n2 < n3:
		fmt.Println("Seu menor número é:", n2)
	case n3 < n1 && n3 < n2:
		fmt.Println("Seu menor número é:", n3)
	}
}
