package main

import (
	"fmt"
	"sync"
)

func sender(messages chan<- string) {
	messages <- "Hello" // Send a message to the messages channel
}

func receiver(messages <-chan string, done chan<- bool) {
	msg := <-messages // Receive a message from the messages channel
	fmt.Println("Received message:", msg)
	done <- true // Signal that the message has been received
}

func main() {
	messages := make(chan string) // Channel for sending messages
	done := make(chan bool)       // Channel for signaling message received

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		sender(messages)
	}()

	go func() {
		defer wg.Done()
		receiver(messages, done)
	}()

	go func() {
		defer wg.Done()
		receiver(messages, done)
	}()

	// Wait for both receivers to finish
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait until both messages are received
	for i := 0; i < 2; i++ {
		<-done
	}

	fmt.Println("All messages received")
}