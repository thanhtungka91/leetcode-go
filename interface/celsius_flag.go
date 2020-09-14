package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {

	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return nil
}

func (f *celsiusFlag) String() string {
	return fmt.Sprintf("%f", f.Celsius)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(f * 5 / 9)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}

	flag.CommandLine.Var(&f, name, usage) // this is the most important, parase flag from command line

	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temparature")

var simpleString = flag.String("test", "test", "test")

func main() {

	flag.Parse()

	fmt.Println(*temp)

	fmt.Println(*simpleString)
}

// similar this one
//https://play.golang.org/p/Ri4qKYTMYbQ

// go build celsius_flag.go
// ./celsius_flag -temp 20F -test hehehehehehe
