package main

import "fmt"

func main() {
	var n1 int
	fmt.Println("Digite a sua idade:")
	fmt.Scan(&n1)
	switch {
	case n1 < 10:
		fmt.Println("Sua classificação é mirim.")
	case n1 > 9 && n1 < 14:
		fmt.Println("Sua classificação é Infantil.")
	case n1 > 13 && n1 < 18:
		fmt.Println("Sua classificação é Juvenil.")
	case n1 > 18:
		fmt.Println("Sua classificação é Adulto.")
	}
}
