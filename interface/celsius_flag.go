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
	return "pk"
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(f * 5 / 9)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}

	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temparature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
