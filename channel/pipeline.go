package main

import "math/rand"

var CloseA = false
var data = make(map[int] bool)

func random(min, max int) int  {

	return rand.Intn(max-min) + min

}

func sendtoOut(min, max int, out chan <-int)  {
	for {
		if CloseA { // close channel
			close(out)
			return
		}

		out <-random(min, max)
	}
}

func reciveddatafromIn(out chan <- int, in <-chan int )  {
	for x := range in {

	}
}