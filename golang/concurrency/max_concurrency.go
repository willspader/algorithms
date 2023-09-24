package main

import (
	"fmt"
	"sync"
	"time"
)

func processItem(item int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the worker as done when finished processing

	// Simulate some processing time
	time.Sleep(time.Millisecond * 500)

	fmt.Printf("Processed item %d\n", item)
}

func main() {
	itemCount := 100
	maxConcurrency := 25

	items := make([]int, itemCount)
	for i := 0; i < itemCount; i++ {
		items[i] = i + 1
	}

	var wg sync.WaitGroup
	concurrencySem := make(chan struct{}, maxConcurrency) // Use a semaphore to limit concurrency

	for _, item := range items {
		concurrencySem <- struct{}{} // Acquire a semaphore slot, limiting concurrency
		wg.Add(1)                    // Increment the WaitGroup counter

		go func(item int) {
			defer func() {
				<-concurrencySem // Release a semaphore slot
			}()

			processItem(item, &wg)
		}(item)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All items processed")
}
