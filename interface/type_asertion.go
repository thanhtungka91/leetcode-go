package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Abser interface {
	Abs() float64
	TestOK() float64
}

type MyFloat float64

func (m MyFloat) Abs() float64 {
	return float64(m)
}

func (m MyFloat) TestOK() float64 {
	return float64(m) * 0.12
}

func (m MyFloat) JustAfunction() string {
	return "hello"
}
func main() {

	var w io.Writer

	fmt.Printf("%T\n", w)
	w = os.Stdout
	f := w.(*os.File)
	// c := w.(*bytes.Buffer)
	fmt.Printf("%T\n", f)
	// fmt.Printf("%T\n", c)

	var i interface{} = "hello"

	s := i.(string)

	fmt.Println(s)

	fmt.Printf("%T\n", i)

	var a Abser

	xx := MyFloat(-math.Sqrt2)

	a = xx

	tttt := a.TestOK()
	fmt.Println("ttttt", tttt)

	fmt.Printf("type of T is %T\n", tttt)
	fmt.Printf("%T\n", a)

}
