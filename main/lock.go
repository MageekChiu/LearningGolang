package main

import (
	"sync"
	"fmt"
)

var (
	mu sync.Mutex
	w sync.WaitGroup
	number = 100
	//ch = make(chan int)
)

func main()  {

	for i:=0;i<1000 ;i++  {
		w.Add(2)
		go func() {
			mu.Lock()
			number += 1
			mu.Unlock()
			w.Done()
			//ch <- 1

		}()

		go func() {
			mu.Lock()
			number -= 1
			mu.Unlock()
			w.Done()
			//ch <- 1

		}()
	}

	//for i:=0;i<1000 ;i++  {
		//<- ch
		//<- ch
	//}

	w.Wait()

	fmt.Println(number)
}



