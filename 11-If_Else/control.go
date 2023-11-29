package main

import "fmt"

func main() {

	numero := 153

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 15 {
		fmt.Println("Maior que 0")
	}
}
