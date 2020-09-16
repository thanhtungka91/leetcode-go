package main

import "fmt"

type square struct {
	X float64
}

func (s square) Perimeter() float64 {
	return s.X
}

type circle struct {
	R float64
}

func (s circle) Area() float64 {
	return 3.14 * s.R
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (s square) Area() float64{
	return s.X*s.X
}

func (s circle) Perimeter() float64 {
	return s.R*2
}

func interfacFunction(shape Shape){
	fmt.Println(shape.Area())
}

// n
func main()  {
	x := square{
		10,
	}
	 y :=circle{
	 	100,
	 }
	interfacFunction(x)
	interfacFunction(y)


}