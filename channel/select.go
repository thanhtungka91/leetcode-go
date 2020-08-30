package main

import "fmt"

func main() {
	message := make(chan string)

	signal := make(chan bool)

	go func() {
		message <- "hi"
	}()

	go func() {
		signal <- false
	}()


	for i := 0; i <= 2; i++{
		select {

		case msg := <-message:
			fmt.Println("received message", msg)
		case sig := <-signal:
			fmt.Println("received signal", sig)

		default:
			fmt.Println("no activity")
		}
	}




}