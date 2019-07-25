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
		stringlengh int
		err error
	}
	done := make(chan Result)
	for _, f := range filenames{
		fmt.Println("before go")
		go func(f string) {
			createDelay(f)
			done <- Result{
				stringlengh:len(f),
				err:nil,
			}


		}(f)

	}

	for range filenames{

		result := <- done

		fmt.Println("wait at least one", result)
		if result.stringlengh == 0 {
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

