package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, value := range numeros {
		total += value
	}
	return total
}

func main() {
	fmt.Println(soma(3, 45, 12, 500))
}
