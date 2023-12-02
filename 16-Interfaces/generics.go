package main

import "fmt"

// Qualquer tipo atende o parametro interf
func generics(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generics("String")
	generics(1)
	generics(true)
	generics(nil)
}
