package main

import "fmt"

func closure(input string) func() {
	text := fmt.Sprintf("Dentro da %s", input)

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {
	text := "Dentro de main"
	fmt.Println(text)

	funcao1 := closure("um")
	funcao2 := closure("dois")
	funcao1()
	funcao2()
}
