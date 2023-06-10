package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)
	time.Sleep(50 * time.Second)
	close(ch)
}

func producer(ch chan<- int) {
	i := 0
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		ch <- i
		i++
		fmt.Printf("producer data = %d\n", i)
	}
}

func consumer(ch <-chan int) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Printf("consumer data = %d\n", <-ch)
	}
}
