package main

import "fmt"

type square struct {
	X float64
}

func (s square) Perimeter() float64 {
	return s.X
}

func (s square) Area() float64 {
	return s.X * s.X
}

type circle struct {
	R float64
}

func (c circle) Area() float64 {
	return 3.14 * c.R
}

func (c circle) Perimeter() float64 {
	return c.R * 2
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func interfaceFunction(sh Shape) {

	_, ok := sh.(circle)
	if ok {
		fmt.Println("this is a circle")
	}

	v, ok := sh.(square)
	if ok {
		fmt.Println("this is a square", v)
	}

	fmt.Println(sh.Area())
}

// n
func main() {
	x := square{
		10,
	}
	y := circle{
		100,
	}

	interfaceFunction(x)
	fmt.Println("===== circle ====")
	interfaceFunction(y)

}
