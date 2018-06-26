package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	ch := make(chan int,10)


	go b(ch)
	go a(ch)
	go b(ch)
	go a(ch)
	a(ch)


	for i:=0;i<5;i++{
		x := <- ch
		fmt.Println(x)
	}


}

func a(ch chan int)  {
	fmt.Println("aaaaaaaaaa--"+time.Now().Format("20060102150405"))
	ch <- rand.Intn(1000)
}

func b(ch chan int)  {
	fmt.Println("bbbbbbbbbb--"+time.Now().Format("20060102150405"))
	ch <- rand.Intn(5)
}