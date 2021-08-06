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
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Line interface {
	Scale(f float64)
}

func fromTo(l Line) {
	switch l.(type) {
	case Vertex:
		fmt.Println("From ", l.(Vertex).X, " -> To ", l.(Vertex).Y)
	case *Vertex:
		fmt.Println("From ", l.(*Vertex).X, " -> To ", l.(*Vertex).Y)
	default:
		panic("Oops type doesn't match any.")
	}
}

func main() {
	v := Vertex{3, 4}
	(&v).scale(10)
	v.scale(10) // a receiver of type *T will implicitly convert T to *T if possible else error.
	// it is attached to pointer of type Vertex not Vertex itself
	//Hence Scale is of type T and is attached to T
	fmt.Println((&v).Abs())

	fromTo(v)
	fromTo(&v)
}
