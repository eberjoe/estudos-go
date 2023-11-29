package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	p1 := pessoa{"John", "White", 23}

	e1 := estudante{p1, "Medicina"}

	fmt.Println(e1)
}
