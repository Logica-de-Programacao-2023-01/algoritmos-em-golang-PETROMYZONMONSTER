//Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores

package main

import "fmt"

func main() {
	var n1 int
	fmt.Println("Digite o seu número:")
	fmt.Scan(&n1)
	for i := 1; i < n1; i++ {
		if n1%i == 0 {
			fmt.Println(i)
		}
	}
}
