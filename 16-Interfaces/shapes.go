package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	height float64
	width  float64
}

// Implementacao implicita do Go
func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func printArea(s shape) {
	fmt.Printf("A area da forma eh %0.2f\n", s.area())
}

func main() {
	r := rectangle{10, 20}
	printArea(r)

	c := circle{10}
	printArea(c)
}
