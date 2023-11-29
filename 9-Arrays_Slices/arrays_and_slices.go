package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]int
	array1[0] = 40
	fmt.Println(array1)

	slice := []int{23, 4223, 1, 5, 0, 123, -90}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice))
}
