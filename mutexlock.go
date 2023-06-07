package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, mutex *sync.Mutex, results chan<- int) {
	defer wg.Done()

	// Perform some work
	fmt.Printf("Worker %d is starting...\n", id)
	time.Sleep(time.Second) // Simulating work

	// Lock the mutex to access the shared resource
	mutex.Lock()

	// Perform critical section operations
	fmt.Printf("Worker %d is inside the critical section\n", id)
	results <- id // Send the result through the channel

	// Unlock the mutex to release the shared resource
	mutex.Unlock()
	fmt.Printf("Worker %d is exiting...\n", id)
}

func main() {
	const numWorkers = 5

	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	results := make(chan int)

	// Create goroutines for the workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, mutex, results)
	}

	// Wait for all the workers to finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Collect the results from the channel
	for result := range results {
		fmt.Printf("Received result from worker %d\n", result)
	}
}