package main

import "fmt"

// Multiple returns
func calculos(n1, n2 int16) (int16, int16) {
	soma := n1 + n2
	produto := n1 * n2

	return soma, produto
}

func main() {
	resultadoSoma, _ := calculos(20, 10)

	fmt.Println(resultadoSoma)
	fmt.Println(calculos(20, 60))
}
