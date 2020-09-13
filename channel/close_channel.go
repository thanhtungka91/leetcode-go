package main

import "fmt"

func main() {
	jobs:= make(chan int)
	done := make(chan bool)

	go func(){ // receive values from jobs channel
		for {
			_, more:= <-jobs

			if more {
				fmt.Println("received job")
			} else {
				fmt.Println("received all")
				done<- true
			}
		}
	}()


	go func() {
		for i := 0 ; i < 10; i ++ {
			jobs <- i  // send value to channel jobs
		}
		// close after send to channel otherwise receive channel will be deadlock
		close(jobs)
	}()


	fmt.Println("sent all jobs")

	<-done
}
