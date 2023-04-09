package main

import "fmt"

func main() {
	for i := 10; i < 11; i-- {
		fmt.Println(i)
		if i == 0 {
			break
		}
	}
}
