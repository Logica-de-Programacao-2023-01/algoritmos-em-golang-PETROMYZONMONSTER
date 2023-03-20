//Faça um algoritmo que imprima os números ímpares de 1 a 19.

package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
