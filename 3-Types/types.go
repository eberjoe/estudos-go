package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -100000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// INT32 = RUNE
	var numero3 rune = 1234567
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 255
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234500293932.67
	fmt.Println(numeroReal2)

	numeroReal3 := 3124.53
	fmt.Println(numeroReal3)

	var str = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'E'
	fmt.Println(char)

	var zerao error = errors.New("Erro doido")
	fmt.Println(zerao)

	var sim bool = true
	fmt.Println(sim)
}
