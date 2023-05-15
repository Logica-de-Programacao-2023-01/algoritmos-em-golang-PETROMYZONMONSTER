package main

import "fmt"

func main() {
	var n1, maior int
	fmt.Println("Digite os números. A leitura irá ser interrompida quando 0 for lido.")
	fmt.Scan(&n1)
	for n1 != 0 {
		fmt.Println("Digite o número:")
		fmt.Scan(&n1)
		if n1 > maior {
			maior = n1
		} else if n1 == 0 {
			fmt.Println("O maior número é:", maior)
		}
	}
}
