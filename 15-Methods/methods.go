package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// a estrutura onde sera inserida o metodo vem antes do nome da func e entre parenteses

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u *usuario) aniversariar() {
	u.idade++
}

func main() {
	usuario1 := usuario{"User 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 34}
	usuario2.salvar()

	usuario1.aniversariar()
	fmt.Println(usuario1)
}
