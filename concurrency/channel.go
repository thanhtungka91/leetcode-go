package main

import (
"fmt"
"time"
)

func main()  {
	list :=[]string{"23232323232","2332323","1113", "12", "1"}
	makeThumbai(list)
}

func makeThumbai(filenames []string){
	type Result struct {
		stringleng int
		err error
	}
	done := make(chan Result)
	for _, f := range filenames{
		fmt.Println("before go")
		go func(f string) {
			createDelay(f)
			fmt.Println("hello", f)
			if len(f) == 2 {
				done <- Result{
					stringleng:len(f),
					err:nil,
				}
			}

		}(f)

	}

	for range filenames{
		fmt.Println("when=====")


		result := <- done

		fmt.Println("wait at least one", result)
		if result.stringleng == 0 {
			break
		}

		time.Sleep(time.Duration(4*time.Second))
	}

	fmt.Println("done")

	//time.Sleep(2*time.Second)
}

func createDelay(str string){

	time.Sleep(time.Duration(len(str))*time.Second)
}

