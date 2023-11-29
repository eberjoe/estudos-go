package main

import "fmt"

func handleRef(ref *int) {
	*ref /= 2
}

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	// Valores diferentes
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	variavel3 = 200
	fmt.Println(variavel3, ponteiro, *ponteiro)

	r := 10
	handleRef(&r)
	fmt.Println(r)
}
