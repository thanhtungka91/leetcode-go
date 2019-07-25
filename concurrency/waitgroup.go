package main

import (
"fmt"
	"sync"
	"time"
)

func main()  {
	list :=[]string{"23232323232","2332323","1113", "12", "1"}
	makeThumbai2(list)
}

func makeThumbai2(filenames []string){
	type Result struct {
		stringlengh int
		err error
	}

	var wg sync.WaitGroup

	done := make(chan Result)

	for _, f := range filenames{
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			createDelay2(f)
			if len(f) != 2 {
				fmt.Println("add to channel", f)
				done <- Result{
					stringlengh:len(f),
					err:nil,
				}
			}

		}(f)

	}

	go func(){
		wg.Wait()
		fmt.Println(len(done))
		close(done)
	}()

	for range done{
		//result := <- done

		fmt.Println("wait at least one", <-done)

		time.Sleep(time.Duration(4*time.Second))
	}

	fmt.Println("done")


}

func createDelay2(str string){

	time.Sleep(time.Duration(len(str))*time.Second)
}

