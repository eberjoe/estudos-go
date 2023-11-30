package main

import "fmt"

func init() { // Executada antes de tudo
	fmt.Println("Executando a funcao init")
}

func main() {
	fmt.Println("Funcao main sendo executada")
}
