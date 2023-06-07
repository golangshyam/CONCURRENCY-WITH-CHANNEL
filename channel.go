package main

import (
	"fmt"
	"sync"
)

func Airhost(comm chan string, wg *sync.WaitGroup) {

	fmt.Println("terminal1")

	comm <- "depature closed"

	fmt.Println("Airhost before done")

	wg.Done()

	fmt.Println("Airhost after done")

}

func Gate(comm chan string, wg *sync.WaitGroup) {

	fmt.Println("terminal2", <-c)

	fmt.Println("gate before done")

	fmt.Println("gate after done")
	wg.Done()
}

func main() {

	var comm chan string = make(chan string) // heap memory
	// write or read os is informed

	var waits sync.WaitGroup

	go Airhost(comm, &waits)

	go Gate(comm, &waits)

	fmt.Println("two goroutines registered ")

	waits.Add(2)

	fmt.Println("waiting for two goroutines  ")

	waits.Wait()

}
