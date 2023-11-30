package main

import "fmt"

func recuperaExec() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperaExec()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!") // KILLS EXECUTION
}

func main() {
	fmt.Println(alunoEstaAprovado(10, 2))
	fmt.Println("Pos execucao")
}
