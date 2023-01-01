package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 10)
	go prod(c)
	go consume(c)
	time.Sleep(time.Minute)
}

func prod(ch chan<- int) {
	fmt.Println("prod start")
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

func consume(ch <-chan int) {
	fmt.Println("consume start")
	for {
		x := <-ch
		fmt.Println(x)
		time.Sleep(time.Second)
	}
}
