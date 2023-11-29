package main

import "fmt"

func main() {
	mapa := map[string]string{
		"Brazil":    "Brasilia",
		"Argentina": "Buenos Aires",
	}
	fmt.Println(mapa)

	mapa["Paraguay"] = "Assumpcion"
	fmt.Println(mapa)

	delete(mapa, "Brazil")
	fmt.Println(mapa)
}
