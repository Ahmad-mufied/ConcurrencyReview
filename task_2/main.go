package main

import (
	"fmt"
	"strings"
	"sync"
)

// TODO: Task 2 - Waiting for Goroutines
// 2.1. Recognize the problem when the main function doesn't wait for the goroutines to finish.
// 2.2. Use the sync.WaitGroup to make sure your main function waits for the goroutines to complete before exiting

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(&wg)
	go printLetter(&wg)

	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	wg.Done()

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetter(wg *sync.WaitGroup) {
	wg.Done()

	// declaring the variable which is an ASCII value of A
	var startingASCIINumber int = 65

	// printing the alphabet from a to j
	for i := 0; i < 10; i++ {
		letter := string(rune(startingASCIINumber + i))

		fmt.Println(strings.ToLower(letter))
	}
}
