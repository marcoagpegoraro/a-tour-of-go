package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, scale float64) {
	v.X = v.X * scale
	v.Y = v.Y * scale
}

func main() {
	vertex := Vertex{2, 3}

	fmt.Println(Abs(vertex))

	Scale(&vertex, 2)

	fmt.Println(vertex)

}
