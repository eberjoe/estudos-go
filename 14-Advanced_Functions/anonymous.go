package main

import "fmt"

func main() {

	fmt.Println(func(text string) string {
		return fmt.Sprintf("Received: %s", text)
	}("parameter"))
}
