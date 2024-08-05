package main

import (
	"fmt"
	"strings"
	"time"
)

// TODO: 1 - Basic Goroutine
// 1.1. Write a function printNumbers that prints numbers from 1 to 10.
// 1.2. Write a function printLetters that prints letters from 'a' to 'j'.
// 1.3. Use goroutines to concurrently run both functions.

func main() {
	go printNumbers()
	go printLetter()

	time.Sleep(100 * time.Millisecond)
}

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetter() {
	// declaring the variable which is an ASCII value of A
	var startingASCIINumber int = 65

	// printing the alphabet from a to j
	for i := 0; i < 10; i++ {
		letter := string(rune(startingASCIINumber + i))

		fmt.Println(strings.ToLower(letter))
	}
}
