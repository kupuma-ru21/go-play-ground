package main

import "fmt"

type Square struct {
	side float32
}

type Triangle struct {
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Area() float32
	Perimeter() float32
}

func main() {
	aI := AreaInterface(&Triangle{base: 10, height: 5})
	fmt.Printf("The triangle has area: %f\n", aI.Area())

	pI := PeriInterface(&Square{side: 5})
	fmt.Printf("The square has area: %f\n", pI.Area())
	fmt.Printf("The square has perimeter: %f\n", pI.Perimeter())
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 { // implement method called on square to calculate its perimeter
	return sq.side * 4
}

func (tr *Triangle) Area() float32 { // implement method called on triangle to calculate its area
	return (tr.base * tr.height) / 2
}
