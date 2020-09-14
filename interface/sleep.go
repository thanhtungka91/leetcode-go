package main

import (
	"flag"
	"fmt"
	"time"
)

// use Duration interface
var period = flag.Duration("period", 1*time.Second, "slee period")

func main() {
	flag.Parse()
	fmt.Println("sleep for a second", *period)

	time.Sleep(*period)

	fmt.Println("done")
}
