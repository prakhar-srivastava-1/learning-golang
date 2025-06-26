package main

import (
	"fmt"
	"sync"
)

// wait group creates a group for threads
var wg sync.WaitGroup

func main() {
	fmt.Println("Starting the main thread")

	// add 1 whenever you run a goroutine
	wg.Add(1)
	// anonymous function
	go func() {
		fmt.Println("Printing from first goroutine")
		// call done to reduce the goroutine count in wg
		wg.Done()
	}()

	// add 1 whenever you run a goroutine
	wg.Add(1)
	// anonymous function
	go func() {
		fmt.Println("Printing from second goroutine")
		// call done to reduce the goroutine count in wg
		wg.Done()
	}()

	fmt.Println("Printing from middle of main thread")

	// add 1 whenever you run a goroutine
	wg.Add(1)
	// anonymous function
	go func() {
		fmt.Println("Printing from third goroutine")
		// call done to reduce the goroutine count in wg
		wg.Done()
	}()

	// add 1 whenever you run a goroutine
	wg.Add(1)
	// anonymous function
	go func() {
		fmt.Println("Printing from fourth goroutine")
		// call done to reduce the goroutine count in wg
		wg.Done()
	}()

	fmt.Println("End of main thread")

	// makes the program wait for all goroutines to complete
	wg.Wait()
}
