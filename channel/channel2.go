package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i:= 0 ; i< 10 ; i++ {
			c <-i
		}
		close(c)
	}()

	//go func() {
	//	for {
	//		fmt.Println(<-c)
	//	}
	//}()


	for i := 0 ; i <  10 ; i ++ {
		fmt.Println(<-c)
	}
	//for n := range c {
	//
	//	fmt.Println(n)
	//}

	//time.Sleep(2*time.Second)
	//fmt.Println("after time out")
}
