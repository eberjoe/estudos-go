package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var eber usuario
	eber.nome = "Eber"
	eber.idade = 23

	eunice := usuario{"Eunice", 19}

	elisa := usuario{nome: "Elisa"}

	fmt.Println(eber, eunice, elisa)
}
