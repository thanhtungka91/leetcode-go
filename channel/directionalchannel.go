package main

import "fmt"

func  main(){
	fmt.Println("test")

	c := make(chan int)

	for j := 0; j <10; j++ {
		go func() {
			for i  := 0 ; i < 10 ; i++ {
				c <- i
			}
			//close(c)
		}()

	}


	for k := 0; k< 90; k++  {
		fmt.Println(k, <-c)
	}
}