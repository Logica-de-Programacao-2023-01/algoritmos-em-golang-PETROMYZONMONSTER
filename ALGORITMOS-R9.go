package main

import "fmt"

func main() {
	var numero, soma, divisor float64
	fmt.Println("Digite sua sequência de números.O número final sera 0.")
	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}
		soma += numero
		divisor++
	}
	if numero == 0 {
		fmt.Println("Sua média é:", soma/divisor)
	}
}
