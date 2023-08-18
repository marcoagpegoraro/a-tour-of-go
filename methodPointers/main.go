package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	vertex := Vertex{2, 3}

	fmt.Println(vertex.Abs())

	vertex.scale(2)
	fmt.Println(vertex)
}
