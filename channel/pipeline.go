package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CloseA = false
var data = make(map[int] bool)

func random(min, max int) int  {

	return rand.Intn(max-min) + min

}

func main()  {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Println( "n1 should be smaller than n2")
		return
	}

	rand.Seed(time.Now().UnixNano())
	ChanelOne  := make(chan int)
	ChanelTwo := make(chan int)

	go sentDataToReceivedChannel(n1,n2,ChanelOne)
	go receivedDataFromInAndSendToOut(ChanelTwo, ChanelOne)
	readingDataFromIn(ChanelTwo)
	fmt.Println("done!!!!")
	fmt.Println(data)
}

// out -> received,
func sentDataToReceivedChannel(min, max int , out chan <- int) {

	for {
		if CloseA {
			close(out)
			return
		}

		out <- random(min, max) // sent value to out channel


	}
}

func readingDataFromIn( in <- chan int)  {
	var sum int
	sum = 0

	for x2 := range in {

		sum = sum + x2
	}

	fmt.Println("summm is %d\n", sum)
}


func receivedDataFromInAndSendToOut(out chan <- int, in <-chan int )  {
	for x := range in {
		fmt.Println("x: ", x)
		_, ok := data[x]
		if ok  {
			fmt.Println("close channel")
			CloseA = true
		} else {
			data[x] = true
			out <-x
		}
	}

	close(out)
}