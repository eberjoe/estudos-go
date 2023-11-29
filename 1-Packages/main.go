package main

import (
	"1stModule/auxiliary"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello Eber")
	auxiliary.Write()

	erro := checkmail.ValidateFormat("12@3.com")
	fmt.Println(erro)
}
