package main

import "fmt"

func main()  {

	// create a new channel
	c := make(chan int)

	go func() {
		c <-500
	}()

	fmt.Println(<-c)



	fmt.Println("exitc")
}

func foo (c chan<- int){
	c <-42
}

func bar(c <-chan  int){
	fmt.Println(<-c)
}