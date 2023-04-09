package main

import "fmt"

func main() {
	var n1 int
	fmt.Println("Digite o seu número:")
	fmt.Scan(&n1)
	for i := 1; i < 11; i++ {
		fmt.Println(n1 * i)
	}
}
