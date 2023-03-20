//Faça um algoritmo que imprima os números pares de 0 a 20.

package main

import "fmt"

func main() {
	for i := 0; i < 21; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
