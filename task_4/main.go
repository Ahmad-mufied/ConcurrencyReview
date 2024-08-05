package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO : Task 4 -  Buffered Channels
// 4.1. Modify the channel in Task 3 to have a buffer of size 5.
// 4.2. Discuss or note down the behavior difference when using a buffered channel versus an unbuffered one.

// Using Buffered Channels

func main() {
	c := make(chan int, 5) // Buffered channel with size 5
	var wg sync.WaitGroup

	wg.Add(2)

	go produceBuffered(c, &wg)
	go consumeBuffered(c, &wg)

	wg.Wait()
}

func produceBuffered(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		fmt.Printf("Produced: %d (buffer size: %d)\n", i, len(c))
		c <- i
	}
	close(c)
}

func consumeBuffered(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range c {
		fmt.Printf("Consumed: %d (buffer size: %d)\n", num, len(c))
		time.Sleep(100 * time.Millisecond) // Simulate some work with the consumed value
	}
}
