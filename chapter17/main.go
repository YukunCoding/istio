package main

import(
	"fmt"
	"time"
)

func main() {
	x := make(chan int, 0)
	go produce(x)
	go consume(x)
	time.Sleep(time.Minute)
}

// prouce 
func produce(ch chan<-int) {
	fmt.Println("Good")
	for i := 50; i >= 0; i-- {
		ch <- i
	}
}

// consume
func consume(ch <-chan int) {
	fmt.Println("Consume")
	for {
		x := <-ch
		fmt.Println(x)
	}
}
