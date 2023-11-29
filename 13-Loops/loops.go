package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
	}

	fmt.Println(i)

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	mapa := map[string]string{"Brasil": "Brasilia", "Australia": "Camberra"}

	for key, value := range mapa {
		fmt.Println(key, value)
	}

	array := []string{"me", "you", "us"}

	for index, value := range array {
		fmt.Println(index, value)
	}

	for _, value := range "Adeus pequeno principe" {
		fmt.Println(value)
	}

	for {
		fmt.Println("Loop infinito")
	}
}
